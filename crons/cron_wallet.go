package crons

import (
	"log"
	"time"

	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/services"
	"github.com/constant-money/constant-web-api/models"
)

const (
	DeltaTimeInSeconds = 60
)

// WalletCron : struct
type WalletCron struct {
	ud        *daos.UserDAO
	walletSrv *services.WalletService
	wallets   map[string]int64
	Running   bool
}

// InitWalletCron :
func InitWalletCron(ud *daos.UserDAO, walletSrv *services.WalletService) *WalletCron {
	return &WalletCron{
		ud:        ud,
		walletSrv: walletSrv,
		wallets:   make(map[string]int64),
	}
}

// ScanWallets : ...
func (wc *WalletCron) ScanWallets() {
	userWallets, _ := wc.ud.AllUserWallets("import_constant")
	for i := 0; i < len(userWallets); i++ {
		uw := userWallets[i]
		balance, err := wc.walletSrv.ScanBalanceOf(uw)
		if err != nil {
			log.Println(err.Error())
		} else {
			if balance.Int64() > 0 && wc.isAbleToFireHook(uw) {
				wc.walletSrv.SendUserWalletHook(uw, balance.Int64())
			}
		}
	}
}

func (wc *WalletCron) isAbleToFireHook(uw *models.UserWallet) (result bool) {
	start := wc.wallets[uw.WalletAddress]
	now := time.Now().UTC().UnixNano() / 1000000000
	if now-start >= DeltaTimeInSeconds {
		wc.wallets[uw.WalletAddress] = now
		result = true
	}
	return
}
