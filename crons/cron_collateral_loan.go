package crons

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/ethereum"
	"github.com/constant-money/constant-event/services"
	helpers "github.com/constant-money/constant-web-api/helpers"
	wm "github.com/constant-money/constant-web-api/models"
	ws "github.com/constant-money/constant-web-api/serializers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CollateralLoanDowntrendRemindLimit uint64 = 80 * 100 // 80%
	CollateralLoanDowntrendLimit       uint64 = 70 * 100 // 70%
)

// CollateralLoan :
type CollateralLoan struct {
	IsRunningAmount                bool
	IsRunningRemind                bool
	IsRunningPayingInterest        bool
	IsRunningPayingInterestOverdue bool
	IsRunningDowntrend             bool
	IsRunningCollect               bool
	LastIndex                      uint
	collateralLoanDAO              *daos.CollateralLoanDAO
	collateralDAO                  *daos.CollateralDAO
	btcClient                      *services.BitcoinService
	conf                           *config.Config
	etherSrv                       *ethereum.Ethereum
}

// NewCollateralLoan :
func NewCollateralLoan(collateralLoanDAO *daos.CollateralLoanDAO, collateralDAO *daos.CollateralDAO, btcClient *services.BitcoinService, etherSrv *ethereum.Ethereum, conf *config.Config) (cl CollateralLoan) {
	cl = CollateralLoan{
		collateralLoanDAO: collateralLoanDAO,
		collateralDAO:     collateralDAO,
		btcClient:         btcClient,
		conf:              conf,
		etherSrv:          etherSrv,
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
		var ids []uint

		for _, collateralLoan := range collateralLoans {
			cl.LastIndex = collateralLoan.ID
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

			collateralAmount := new(big.Float).Quo(new(big.Float).SetUint64(collateralLoan.CollateralAmount), new(big.Float).SetUint64(wm.AmountDenominator))
			compareValue := new(big.Float).Quo(addrValue, collateralAmount)
			comparePercent := new(big.Float)
			comparePercent.SetString("0.9999") //99,99%

			if compareValue.Cmp(comparePercent) >= 0 {
				ids = append(ids, collateralLoan.ID)
			}

			if ids != nil && len(ids) > 0 {
				cl.sendToHook(ids, ws.CollateralLoanActionWallet)
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

// ScanCollateralPayingInterestOverdue :
func (cl *CollateralLoan) ScanCollateralPayingInterestOverdue() {
	var (
		limit = 1
		page  = 0
	)
	for {
		page++
		collateralLoans, err := cl.collateralLoanDAO.FindAllPayingInterestByDay(3, page, limit)
		if err != nil {
			log.Println("FindAllPayingOnDateOverdue error", err.Error())
			return
		}

		if len(collateralLoans) == 0 {
			return
		}

		var ids []uint
		for _, collateralLoan := range collateralLoans {
			ids = append(ids, collateralLoan.ID)
		}

		if ids != nil && len(ids) > 0 {
			cl.sendToHook(ids, ws.CollateralLoanActionOverdue)
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
		collateralLoans, err := cl.collateralLoanDAO.FindAllPayingInterestByDay(1, page, limit)
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

		if ids != nil && len(ids) > 0 {
			cl.sendToHook(ids, ws.CollateralLoanActionExpired)
		}
	}
}

// ScanCollateralDowntrend :
func (cl *CollateralLoan) ScanCollateralDowntrend() {
	// cl.remindByDate(1)

	collaterals, err := cl.collateralDAO.FindAll()
	if err != nil {
		log.Println("ScanCollateralDowntrend error", err.Error())
		return
	}

	for _, collateral := range collaterals {
		cl.remindDownTrend(collateral.Value / CollateralLoanDowntrendRemindLimit)
		cl.paymentDownTrend(collateral.Value / CollateralLoanDowntrendLimit)
	}
}

// ScanCollateralCollect :
func (cl *CollateralLoan) ScanCollateralCollect() {
	var (
		limit = 1
		page  = 0
	)

	if cl.conf.MasterEthWallet == "" || cl.conf.MasterBtcWallet == "" || cl.conf.MasterUsdtWallet == "" {
		log.Println("Master wallet invalid!")
	}

	for {
		page++
		collateralLoans, err := cl.collateralLoanDAO.FindAllCollect(page, limit)
		if err != nil {
			log.Println("FindAllCollect error", err.Error())
			return
		}

		if len(collateralLoans) == 0 {
			return
		}

		for _, collateralLoan := range collateralLoans {
			if strings.ToLower(collateralLoan.Collateral.Symbol) == "eth" {
				tnxHash, err := cl.collectETH(collateralLoan)
				if err != nil {
					continue
				}
				fmt.Println(tnxHash)
			} else if strings.ToLower(collateralLoan.Collateral.Symbol) == "btc" {
				tnxHash, err := cl.collectBTC(collateralLoan)
				if err != nil {
					continue
				}
				fmt.Println(tnxHash)
			} else if strings.ToLower(collateralLoan.Collateral.Symbol) == "usdt" {
				// Transfer usdt
			}
			// add txes "offchain"
		}
	}
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

		if ids != nil && len(ids) > 0 {
			cl.sendToHook(ids, ws.CollateralLoanActionRemind)
		}
	}
}

func (cl *CollateralLoan) remindDownTrend(currentValue uint64) {
	var (
		limit = 1
		page  = 0
	)
	for {
		page++
		collateralLoans, err := cl.collateralLoanDAO.FindAllDowntrend(currentValue, page, limit)
		if err != nil {
			log.Println("FindAllDowntrend error", err.Error())
			break
		}

		if len(collateralLoans) == 0 {
			break
		}

		var ids []uint
		for _, collateralLoan := range collateralLoans {
			ids = append(ids, collateralLoan.ID)
		}

		if ids != nil && len(ids) > 0 {
			cl.sendToHook(ids, ws.CollateralLoanActionDownTrendRemind)
		}
	}
}

func (cl *CollateralLoan) paymentDownTrend(currentValue uint64) {
	var (
		limit = 1
		page  = 0
	)
	for {
		page++
		collateralLoans, err := cl.collateralLoanDAO.FindAllDowntrend(currentValue, page, limit)
		if err != nil {
			log.Println("FindAllDowntrend error", err.Error())
			break
		}

		if len(collateralLoans) == 0 {
			break
		}

		var ids []uint
		for _, collateralLoan := range collateralLoans {
			ids = append(ids, collateralLoan.ID)
		}

		if ids != nil && len(ids) > 0 {
			cl.sendToHook(ids, ws.CollateralLoanActionDownTrend)
		}
		// TODO sell coin
	}
}

func (cl *CollateralLoan) sendToHook(ids []uint, action ws.CollateralLoanAction) {
	jsonWebhook := make(map[string]interface{})
	jsonWebhook["type"] = ws.WebhookTypeCollateralLoan
	jsonWebhook["data"] = map[string]interface{}{
		"Action": action,
		"IDs":    ids,
	}

	err := hookService.Event(jsonWebhook)
	if err != nil {
		log.Println("Hook remind success error: ", jsonWebhook, err.Error())
	}
}

func (cl *CollateralLoan) collectETH(collateralLoan *wm.CollateralLoan) (string, error) {
	priKey, err := helpers.DecryptToString(collateralLoan.CollateralPrivateKey, cl.conf.CipherKey)
	if err != nil {
		priKey = collateralLoan.CollateralPrivateKey
	}

	gasPrice, err := cl.etherSrv.GetGasPrice()
	if err != nil {
		log.Println("GetGasPrice error", err.Error())
		return "", err
	}

	balanceAtAddr, err := cl.etherSrv.BalanceAtAddr(collateralLoan.CollateralAddress)
	if err != nil {
		log.Println("BalanceAtAddr error", err.Error())
		return "", err
	}

	gas := new(big.Int)
	gas = gas.SetInt64(300000)

	gasFee := new(big.Int)
	gasFee = gasFee.Mul(gasPrice, gas)

	value := new(big.Int)
	value = value.Sub(balanceAtAddr, gasFee)

	if value.Cmp(big.NewInt(0)) == -1 || value.Cmp(big.NewInt(0)) == 0 {
		log.Println("value invalid: ", value)
		return "", errors.New("value invalid")
	}

	tnxHash, err := cl.etherSrv.SendSignedTransaction(priKey, cl.conf.MasterEthWallet, value, nil)
	if err != nil {
		log.Println("value", value)
		log.Println("SendSignedTransaction error", err.Error())
		return "", err
	}
	return tnxHash, nil
}

func (cl *CollateralLoan) collectBTC(collateralLoan *wm.CollateralLoan) (string, error) {
	result, err := cl.btcClient.BTCSendRawTnx(collateralLoan.CollateralAddress, collateralLoan.CollateralPrivateKey, cl.conf.CipherKey, "mru5KyqSSnG1sgTwXmaqhuGomMkd8TVeGd", 100)
	fmt.Println(result)
	fmt.Println(err)
	return "", nil
}
