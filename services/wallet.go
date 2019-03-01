package services

import (
	"errors"
	"fmt"

	"github.com/constant-money/constant-event/ethereum"
	wm "github.com/constant-money/constant-web-api/models"
)

// WalletService : struct
type WalletService struct {
	constant     *ethereum.Constant
	hookEndpoint string
}

// InitWalletService :
func InitWalletService(constant *ethereum.Constant, hookEndpoint string) *WalletService {
	return &WalletService{
		constant:     constant,
		hookEndpoint: hookEndpoint,
	}
}

// func (us *UserService) sendUserWalletHook(userWalletID uint, walletAddr string, masterAddr string, metaData string) error {
// 	jsonData := make(map[string]interface{})
// 	jsonData["type"] = 3
// 	jsonData["data"] = map[string]interface{}{
// 		"from":     walletAddr,
// 		"to":       masterAddr,
// 		"metaData": metaData,
// 		"id":       userWalletID,
// 	}

// 	endpoint := us.hookEndpoint
// 	endpoint = fmt.Sprintf("%s", endpoint)
// 	jsonValue, _ := json.Marshal(jsonData)

// 	request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
// 	request.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}

// 	b, _ := ioutil.ReadAll(response.Body)

// 	var data map[string]interface{}
// 	json.Unmarshal(b, &data)

// 	status, ok := data["status"]
// 	message, hasMessage := data["message"]

// 	if ok && status.(float64) > 0 {
// 		return nil
// 	} else {
// 		errStr := "Unknown"
// 		if hasMessage {
// 			errStr = message.(string)
// 		}
// 		return errors.New(errStr)
// 	}
// }

// ScanBalanceOf : ...
func (ws *WalletService) ScanBalanceOf(userWallet *wm.UserWallet) error {
	if userWallet != nil {
		walletAddress := userWallet.WalletAddress
		if walletAddress != "" {
			bal, err := ws.constant.BalanceOf(walletAddress)
			if err != nil {
				return err
			}
			fmt.Println("WTF = ", bal.Uint64())
		}
		return nil
	}
	return errors.New("Invalid wallet address")
}
