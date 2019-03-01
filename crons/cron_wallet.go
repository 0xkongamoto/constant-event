package crons

import (
	"fmt"
	"log"

	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/services"
)

// WalletCron : struct
type WalletCron struct {
	ud        *daos.UserDAO
	walletSrv *services.WalletService
	Running   bool
}

// InitWalletCron :
func InitWalletCron(ud *daos.UserDAO, walletSrv *services.WalletService) *WalletCron {
	return &WalletCron{
		ud:        ud,
		walletSrv: walletSrv,
	}
}

// ScanWallets : ...
func (wc *WalletCron) ScanWallets() {
	userWallets, _ := wc.ud.AllUserWallets("import_constant")
	fmt.Println("Wallets = ", len(userWallets))
	for i := 0; i < len(userWallets); i++ {
		uw := userWallets[i]
		err := wc.walletSrv.ScanBalanceOf(uw)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}
