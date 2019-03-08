package crons

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/services"
	wm "github.com/constant-money/constant-web-api/models"
	ws "github.com/constant-money/constant-web-api/serializers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CollateralLoan :
type CollateralLoan struct {
	IsRunningAmount         bool
	IsRunningRemind         bool
	IsRunningPayingInterest bool
	LastIndex               uint
	collateralLoanDAO       *daos.CollateralLoanDAO
	btcClient               *services.BitcoinService
	conf                    *config.Config
}

// NewCollateralLoan :
func NewCollateralLoan(collateralLoanDAO *daos.CollateralLoanDAO, btcClient *services.BitcoinService, conf *config.Config) (cl CollateralLoan) {
	cl = CollateralLoan{
		collateralLoanDAO: collateralLoanDAO,
		btcClient:         btcClient,
		conf:              conf,
	}
	return cl
}

// ScanCollateralAmount :
func (cl *CollateralLoan) ScanCollateralAmount() {
	conf := config.GetConfig()
	networkURL := conf.ChainURL

	collateralLoans, err := cl.collateralLoanDAO.FindAllPending(cl.LastIndex, 10)

	if err != nil {
		log.Println("Find Collateral Loans peding error", err.Error())
		return
	}
	if len(collateralLoans) == 0 {
		cl.LastIndex = 0
	} else {
		cl.LastIndex = collateralLoans[0].ID
		var ids []uint

		for _, collateralLoan := range collateralLoans {
			var (
				balanceStr = ""
				decimal    = 0
			)

			if collateralLoan.Collateral.WalletType == wm.CollateralWalletTypeEthereum {
				account := common.HexToAddress(collateralLoan.CollateralAddress)
				etherClient, err := ethclient.Dial(networkURL)
				balance, err := etherClient.BalanceAt(context.Background(), account, nil)
				if err != nil {
					log.Println(err)
					continue
				}
				balanceStr = balance.String()
				decimal = 18
			} else if collateralLoan.Collateral.WalletType == wm.CollateralWalletTypeBitcoin {
				balance, err := cl.btcClient.BTCBalanceOf(collateralLoan.CollateralAddress)
				if err != nil {
					log.Println(err)
					continue
				}
				balanceStr = balance
				decimal = 10
			} else {
				log.Println("Not found wallet type: ", collateralLoan)
				continue
			}

			fbalance := new(big.Float)
			fbalance.SetString(balanceStr)
			addrValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(decimal)))
			addrValue = new(big.Float).Mul(big.NewFloat(100), addrValue)

			if addrValue.Cmp(new(big.Float).SetUint64(collateralLoan.CollateralAmount)) >= 0 {
				ids = append(ids, collateralLoan.ID)
			}

			jsonWebhook := make(map[string]interface{})
			jsonWebhook["type"] = ws.WebhookTypeCollateralLoan
			jsonWebhook["data"] = map[string]interface{}{
				"Action": ws.CollateralLoanActionWallet,
				"IDs":    ids,
			}

			err = hookService.Event(jsonWebhook)
			if err != nil {
				log.Println("Hook loan wallet success error: ", err.Error())
			}
		}
	}
}

// ScanCollateralRemind :
func (cl *CollateralLoan) ScanCollateralRemind() {
	cl.remindByDate(5)
	cl.remindByDate(3)
	cl.remindByDate(1)
}

func (cl *CollateralLoan) remindByDate(dayNumber uint) {
	var (
		limit = 1
		page  = 0
	)
	for {
		page++
		collateralLoans, err := cl.collateralLoanDAO.FindAllPayingByDate(dayNumber, page, limit)
		if err != nil {
			log.Println("FindAllPayingByDate error", err.Error())
			return
		}

		if len(collateralLoans) == 0 {
			return
		}

		var ids []uint
		for _, collateralLoan := range collateralLoans {
			ids = append(ids, collateralLoan.ID)
		}

		jsonWebhook := make(map[string]interface{})
		jsonWebhook["type"] = ws.WebhookTypeCollateralLoan
		jsonWebhook["data"] = map[string]interface{}{
			"Action": ws.CollateralLoanActionRemind,
			"IDs":    ids,
		}

		err = hookService.Event(jsonWebhook)
		if err != nil {
			log.Println("Hook remind success error: ", err.Error())
		}
	}
}

// ScanCollateralPayingInterest :
func (cl *CollateralLoan) ScanCollateralPayingInterest() {
	var (
		limit = 1
		page  = 0
	)
	for {
		page++
		collateralLoans, err := cl.collateralLoanDAO.FindAllPayingLastDay(page, limit)
		if err != nil {
			log.Println("FindAllPayingOnDate error", err.Error())
			return
		}

		if len(collateralLoans) == 0 {
			return
		}

		var ids []uint
		for _, collateralLoan := range collateralLoans {
			ids = append(ids, collateralLoan.ID)
		}

		jsonWebhook := make(map[string]interface{})
		jsonWebhook["type"] = ws.WebhookTypeCollateralLoan
		jsonWebhook["data"] = map[string]interface{}{
			"Action": ws.CollateralLoanActionExpired,
			"IDs":    ids,
		}

		err = hookService.Event(jsonWebhook)
		if err != nil {
			log.Println("Hook remind success error: ", err.Error())
		}
	}
}
