package crons

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
	wm "github.com/constant-money/constant-web-api/models"
	ws "github.com/constant-money/constant-web-api/serializers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
)

// CollateralLoan :
type CollateralLoan struct {
	IsRunningAmount         bool
	IsRunningRemind         bool
	IsRunningPayingInterest bool
	LastIndex               uint
	collateralLoanDAO       *daos.CollateralLoanDAO
	conf                    *config.Config
}

// NewCollateralLoan :
func NewCollateralLoan(collateralLoanDAO *daos.CollateralLoanDAO, conf *config.Config) (cl CollateralLoan) {
	cl = CollateralLoan{
		collateralLoanDAO: collateralLoanDAO,
		conf:              conf,
	}
	return cl
}

// ScanCollateralAmount :
func (cl *CollateralLoan) ScanCollateralAmount() {
	conf := config.GetConfig()
	networkURL := conf.ChainURL

	etherClient, err := ethclient.Dial(networkURL)

	collateralLoans, err := cl.collateralLoanDAO.FindAllPending(cl.LastIndex, 10)
	if err != nil {
		log.Println("Find Collateral Loans peding error", err.Error())
		return
	}
	if len(collateralLoans) == 0 {
		cl.LastIndex = 0
	} else {
		cl.LastIndex = collateralLoans[0].ID
		for index := 0; index < len(collateralLoans); index++ {
			account := common.HexToAddress(collateralLoans[index].CollateralAddress)
			balance, err := etherClient.BalanceAt(context.Background(), account, nil)
			if err != nil {
				log.Fatal(err)
				return
			}

			fbalance := new(big.Float)
			fbalance.SetString(balance.String())
			ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			ethValue = new(big.Float).Mul(big.NewFloat(100), ethValue)

			if ethValue.Cmp(new(big.Float).SetUint64(collateralLoans[index].CollateralAmount)) >= 0 {
				collateralLoans[index].Status = wm.CollateralLoanStatusAccepted
				errTx := models.WithTransaction(func(tx *gorm.DB) error {
					if err := cl.collateralLoanDAO.Update(tx, collateralLoans[index]); err != nil {
						log.Println("Update Collateral Loan error", err.Error())
						return err
					}
					return nil
				})

				if errTx != nil {
					log.Println("DB Tnx Update Collateral Loan error", errTx.Error())
				}
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

		for _, collateralLoan := range collateralLoans {
			collateralLoan.Status = wm.CollateralLoanStatusPayingInterest
			errTx := models.WithTransaction(func(tx *gorm.DB) error {
				// TODO send to hook
				if err := cl.collateralLoanDAO.Update(tx, collateralLoan); err != nil {
					log.Println("Update Collateral Loan status error", err.Error())
					return err
				}
				return nil
			})

			if errTx != nil {
				log.Println("DB Tnx Update Collateral Loan status error", errTx.Error())
			}
		}
	}
}
