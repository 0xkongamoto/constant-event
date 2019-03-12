package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/crons"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/ethereum"
	"github.com/constant-money/constant-event/models"
	"github.com/constant-money/constant-event/services"
	"github.com/robfig/cron"

	"github.com/constant-money/constant-web-api/services/3rd/primetrust"
)

func main() {
	// config Logger
	logFile, err := os.OpenFile("logs/log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile) // You may need this
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// config
	conf := config.GetConfig()
	pt := primetrust.Init(conf.PrimetrustEndpoint, conf.PrimetrustUsername, conf.PrimetrustPassword, conf.PrimetrustAccountID)

	mapContracts := conf.Contracts
	for _, value := range mapContracts {
		fmt.Printf("Start cron scan contract : %s %s \n", value.Name, value.Address)

		cr := crons.NewCron(value.Name, value.Address, &daos.MasterAddressDAO{})
		appCron := cron.New()
		appCron.AddFunc("@every 15s", func() {
			fmt.Println("scan tx every 15s")
			if !cr.ScanRunning {
				cr.ScanRunning = true
				cr.ScanTx()
				cr.ScanRunning = false
			} else {
				fmt.Println("scan tx is running")
			}
		})
		appCron.AddFunc("@every 15m", func() {
			fmt.Println("sync tx every 15m")
			if !cr.SyncRunning {
				cr.SyncRunning = true
				cr.SyncTx()
				cr.SyncRunning = false
			} else {
				fmt.Println("sync tx is running")
			}
		})

		appCron.Start()
		time.Sleep(time.Second * 1)
	}

	// add order cron
	orderDAO := daos.InitOrderDAO(models.Database())
	orderSrv := services.InitOrderService(orderDAO, conf)
	orderCron := cron.New()
	orderCron.AddFunc("@every 1m", func() {
		fmt.Println("scan order every 1m")
		if !orderSrv.Running {
			orderSrv.Running = true
			orderSrv.ScanOrders()
			orderSrv.Running = false
		} else {
			fmt.Println("scan order is running")
		}
	})
	orderCron.Start()

	// add exchange rate cron
	exchangeDAO := daos.InitExchangeDAO(models.Database())
	exchangeSrv := services.InitExchangeService(exchangeDAO)
	exchangeSrv.ParseExchangeRate()
	exchangeRateCron := cron.New()
	exchangeRateCron.AddFunc("@every 10m", func() {
		fmt.Println("scan exchange rate every 10m")
		if !exchangeSrv.Running {
			exchangeSrv.Running = true
			exchangeSrv.ParseExchangeRate()
			exchangeSrv.Running = false
		} else {
			fmt.Println("scan exchange rate is running")
		}
	})
	exchangeRateCron.Start()

	// add reserve cron
	reserveDAO := daos.NewReserve(models.Database())
	reserveSrv := services.InitReserveService(reserveDAO, pt, conf.HookEndpoint)
	reserveCron := cron.New()
	reserveCron.AddFunc("@every 30m", func() {
		fmt.Println("scan reserve every 30m")
		if !reserveSrv.Running {
			reserveSrv.Running = true
			reserveSrv.PrimetrustHook()
			reserveSrv.Running = false
		} else {
			fmt.Println("scan reserve is running")
		}
	})
	reserveCron.Start()

	// add user wallet cron
	masterAddressDAO := &daos.MasterAddressDAO{}
	userDAO := daos.InitUserDAO(models.Database())
	masterAddr, err := masterAddressDAO.GetMasterAddress()
	etherService := ethereum.Init(conf)

	if err == nil {
		for _, value := range mapContracts {
			constant := ethereum.InitConstant(value.Address, masterAddr.PriKey, conf.CipherKey, etherService)
			walletSrv := services.InitWalletService(constant, conf.HookEndpoint)
			ucWallet := crons.InitWalletCron(userDAO, walletSrv)

			userWalletsCron := cron.New()
			userWalletsCron.AddFunc("@every 2s", func() {
				fmt.Println("scan user wallet service every 2s")
				if !ucWallet.Running {
					ucWallet.Running = true
					ucWallet.ScanWallets()
					ucWallet.Running = false
				} else {
					fmt.Println("scan user wallet service is running")
				}
			})
			userWalletsCron.Start()
			time.Sleep(time.Second * 1)
		}

	} else {
		fmt.Println("cannot start user service cause there is no master address!!!")
	}

	// add user kyc cron
	userSrv := services.InitUserService(pt, conf.HookEndpoint, conf.PrimetrustEndpoint)
	ucKYC := crons.InitUserCron(userDAO, userSrv)
	userKYCCron := cron.New()
	userKYCCron.AddFunc("@every 30m", func() {
		fmt.Println("scan user kyc service every 30m")
		if !ucKYC.Running {
			ucKYC.Running = true
			ucKYC.ScanKYC()
			ucKYC.Running = false
		} else {
			fmt.Println("scan user kyc service is running")
		}
	})
	userKYCCron.Start()

	// add task cron
	crTask := crons.NewCronTask(masterAddressDAO, &daos.TaskDAO{}, &daos.TxDAO{}, etherService, config.GetConfig())
	taskCron := cron.New()
	taskCron.AddFunc("@every 5s", func() {
		fmt.Println("scan task every 5s")
		if !crTask.ScanRunning {
			crTask.ScanRunning = true
			crTask.ScanTask()
			crTask.ScanRunning = false
		} else {
			fmt.Println("scan task is running")
		}
	})
	taskCron.Start()

	// collaterals loan group
	collateralDAO := daos.InitCollateralDAO(models.Database())
	collateralSrv := services.NewCollateralService(models.Database(), collateralDAO)
	collateralCron := cron.New()
	collateralCron.AddFunc("@every 5m", func() {
		fmt.Println("scan collateral rate every 5m")
		if !collateralSrv.RateFeeding {
			collateralSrv.RateFeeding = true
			collateralSrv.RateFeed()
			collateralSrv.RateFeeding = false
		} else {
			fmt.Println("scan collateral rate is running")
		}
	})
	collateralCron.Start()

	btcClientService := services.NewBitcoinService(conf)
	collateralLoanDAO := daos.InitCollateralLoanDAO(models.Database())
	collateralLoanCron := cron.New()
	collateralLoan := crons.NewCollateralLoan(collateralLoanDAO, collateralDAO, btcClientService, conf)
	// Scan ETH wallet
	collateralLoanCron.AddFunc("@every 10s", func() {
		fmt.Println("collateral loan run every 10s")
		if !collateralLoan.IsRunningAmount {
			collateralLoan.IsRunningAmount = true
			collateralLoan.ScanCollateralAmount()
			collateralLoan.IsRunningAmount = false
		} else {
			fmt.Println("collateral loan amount is running")
		}
	})

	// Scan collateral loan remind
	collateralLoanCron.AddFunc("@every 24h", func() {
		fmt.Println("collateral loan run every 24h")
		if !collateralLoan.IsRunningRemind {
			collateralLoan.IsRunningRemind = true
			collateralLoan.ScanCollateralRemind()
			collateralLoan.IsRunningRemind = false
		} else {
			fmt.Println("collateral loan remind is running")
		}

		if !collateralLoan.IsRunningDowntrend {
			collateralLoan.IsRunningDowntrend = true
			collateralLoan.ScanCollateralDowntrend()
			collateralLoan.IsRunningDowntrend = false
		} else {
			fmt.Println("collateral loan remind is running")
		}
	})

	// Scan collateral loan paying interest status
	collateralLoanCron.AddFunc("@every 12h", func() {
		fmt.Println("collateral loan run every 12h")
		if !collateralLoan.IsRunningPayingInterest {
			collateralLoan.IsRunningPayingInterest = true
			collateralLoan.ScanCollateralPayingInterest()
			collateralLoan.IsRunningPayingInterest = false
		} else {
			fmt.Println("collateral loan update status Paying Interest is running")
		}

		if !collateralLoan.IsRunningPayingInterestOverdue {
			collateralLoan.IsRunningPayingInterestOverdue = true
			collateralLoan.ScanCollateralPayingInterestOverdue()
			collateralLoan.IsRunningPayingInterestOverdue = false
		} else {
			fmt.Println("collateral loan update status Paying Interest Overdue is running")
		}
	})
	collateralLoanCron.Start()

	select {}
}
