package crons

import (
	"fmt"
	"log"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/services"
	"github.com/constant-money/constant-web-api/services/3rd/primetrust"
)

// UserCron : struct
type UserCron struct {
	ud         *daos.UserDAO
	conf       *config.Config
	primetrust *primetrust.Primetrust
	userSrv    *services.UserService
	Running    bool
}

// InitUserCron :
func InitUserCron(ud *daos.UserDAO, primetrust *primetrust.Primetrust, userSrv *services.UserService, conf *config.Config) *UserCron {
	return &UserCron{
		ud:         ud,
		primetrust: primetrust,
		userSrv:    userSrv,
		conf:       conf,
	}
}

// ScanKYC : ...
func (uc *UserCron) ScanKYC() {
	users, err := uc.ud.GetAllUsersNeedCheckKYC()
	if err == nil {
		for i := 0; i < len(*users); i++ {
			u := (*users)[i]
			status, errStr := uc.userSrv.CheckPrimetrustContactID(u.PrimetrustContactID)
			if errStr != "404" {
				if status {
					uc.userSrv.SendKYCHook(u.ID, status, errStr)
				} else {
					if u.VerifiedLevel == 4 {
						uc.userSrv.SendKYCHook(u.ID, status, errStr)
					}
				}
			}
		}
	}
}

// ScanWallets : ...
func (uc *UserCron) ScanWallets() {
	userWallets, _ := uc.ud.AllUserWallets("import_constant")
	fmt.Println("Wallets = ", len(userWallets))
	for i := 0; i < len(userWallets); i++ {
		uw := userWallets[i]
		err := uc.userSrv.ScanBalanceOf(uw)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}
