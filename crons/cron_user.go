package crons

import (
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/services"
)

// UserCron : struct
type UserCron struct {
	ud      *daos.UserDAO
	userSrv *services.UserService
	Running bool
}

// InitUserCron :
func InitUserCron(ud *daos.UserDAO, userSrv *services.UserService) *UserCron {
	return &UserCron{
		ud:      ud,
		userSrv: userSrv,
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
