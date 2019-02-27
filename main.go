package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/crons"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
	"github.com/constant-money/constant-event/services"
	"github.com/robfig/cron"

	bedaos "github.com/constant-money/constant-web-api/daos"
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
	userDAO := daos.InitUserDAO(models.Database())
	userSrv := services.InitUserService(userDAO, pt, conf)
	userCron := cron.New()
	userCron.AddFunc("@every 30m", func() {
		fmt.Println("scan user service every 30m")
		if !userSrv.Running {
			userSrv.Running = true
			userSrv.ScanKYC()
			userSrv.ScanWallets()
			userSrv.Running = false
		} else {
			fmt.Println("scan user service is running")
		}
	})
	userCron.Start()

	// add task cron
	crTask := crons.NewCronTask(1, &daos.MasterAddressDAO{}, &daos.TaskDAO{}, &daos.TxDAO{}, config.GetConfig())
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
	collateralDAO := bedaos.NewCollateral()
	collateralSrv := services.NewCollateralService(models.Database(), collateralDAO)
	collateralCron := cron.New()
	collateralCron.AddFunc("@every 30m", func() {
		fmt.Println("scan collateral rate every 30m")
		if !collateralSrv.RateFeeding {
			collateralSrv.RateFeeding = true
			collateralSrv.RateFeed()
			collateralSrv.RateFeeding = false
		} else {
			fmt.Println("scan collateral rate is running")
		}
	})
	collateralCron.Start()

	collateralLoanDAO := daos.InitCollateralLoanDAO(models.Database())
	collateralLoanCron := cron.New()
	collateralLoan := crons.NewCollateralLoan(collateralLoanDAO, conf)
	collateralLoanCron.AddFunc("@every 10s", func() {
		fmt.Println("collateral loan run every 10s")
		if !collateralLoan.IsRunning {
			collateralLoan.IsRunning = true
			collateralLoan.ScanCollateralAmount()
			collateralLoan.IsRunning = false
		} else {
			fmt.Println("collateral loan is running")
		}
	})
	collateralLoanCron.Start()

	select {}
}
