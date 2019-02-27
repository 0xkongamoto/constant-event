package crons

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
)

// CollateralLoan :
type CollateralLoan struct {
	IsRunning         bool
	collateralLoanDAO *daos.CollateralLoanDAO
	conf              *config.Config
}

// NewCollateralLoan :
func NewCollateralLoan(collateralLoanDAO *daos.CollateralLoanDAO, conf *config.Config) (cl CollateralLoan) {
	cl = CollateralLoan{false, collateralLoanDAO, config.GetConfig()}
	return cl
}

// ScanCollateralAmount :
func (cl *CollateralLoan) ScanCollateralAmount() {
	conf := config.GetConfig()
	networkURL := conf.ChainURL

	etherClient, err := ethclient.Dial(networkURL)

	collateralLoans, err := cl.collateralLoanDAO.FindAllPending(10)
	if err != nil {
		log.Println("Find Collateral Loans peding error", err.Error())
		return
	}
	if len(collateralLoans) == 0 {
		log.Println("Collateral Loans: empty")
	} else {

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
				collateralLoans[index].Status = models.CollateralLoanStatusAccepted
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
				return
			}
		}
	}
}
