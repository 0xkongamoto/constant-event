package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	"github.com/constant-money/constant-event/ethereum"
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/constant-money/constant-web-api/serializers"
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

// SendUserWalletHook : ...
func (ws *WalletService) SendUserWalletHook(userWallet *wm.UserWallet, constantAmount int64) error {
	jsonData := make(map[string]interface{})
	jsonData["type"] = serializers.WebhookTypeUserWallet
	jsonData["data"] = map[string]interface{}{
		"user":     userWallet.User.ID,
		"wallet":   userWallet.WalletAddress,
		"source":   userWallet.Source,
		"constant": constantAmount,
	}

	endpoint := ws.hookEndpoint
	endpoint = fmt.Sprintf("%s", endpoint)
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	b, _ := ioutil.ReadAll(response.Body)

	var data map[string]interface{}
	json.Unmarshal(b, &data)

	status, ok := data["status"]
	message, hasMessage := data["message"]

	if ok && status.(float64) > 0 {
		return nil
	}
	errStr := "Unknown"
	if hasMessage {
		errStr = message.(string)
	}
	return errors.New(errStr)
}

// ScanBalanceOf : ...
func (ws *WalletService) ScanBalanceOf(userWallet *wm.UserWallet) (*big.Int, error) {
	if userWallet != nil {
		walletAddress := userWallet.WalletAddress
		if walletAddress != "" {
			bal, err := ws.constant.BalanceOf(walletAddress)
			if err != nil {
				return nil, err
			}
			return bal, nil
		}
	}
	return nil, errors.New("Invalid wallet address")
}
