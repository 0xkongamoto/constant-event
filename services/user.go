package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/constant-money/constant-web/event/config"
	"github.com/constant-money/constant-web/event/daos"
	"github.com/constant-money/constant-web/event/models"
)

// UserService : struct
type UserService struct {
	ud      *daos.UserDAO
	conf    *config.Config
	Running bool
}

// InitUserService :
func InitUserService(userDao *daos.UserDAO, cf *config.Config) *UserService {
	return &UserService{
		ud:   userDao,
		conf: cf,
	}
}

// ScanWallets : ...
func (us *UserService) ScanWallets() {
	userWallets, _ := us.ud.GetAllUserWalletPending()
	for i := 0; i < len(userWallets); i++ {
		uw := userWallets[i]
		err := us.scanTnx(uw.ID, strings.ToLower(uw.WalletAddress), uw.Metadata, uw.ExpiredAt, uw.StartedAt)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (us *UserService) sendUserWalletHook(userWalletID uint, walletAddr string, masterAddr string, metaData string) error {
	jsonData := make(map[string]interface{})
	jsonData["type"] = 3
	jsonData["data"] = map[string]interface{}{
		"from":     walletAddr,
		"to":       masterAddr,
		"metaData": metaData,
		"id":       userWalletID,
	}

	endpoint := us.conf.HookEndpoint
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
	} else {
		errStr := "Unknown"
		if hasMessage {
			errStr = message.(string)
		}
		return errors.New(errStr)
	}
}

func (us *UserService) scanTnx(ID uint, walletAddr string, metaData string, expiredAt int64, startedAt int64) error {
	masterAddrArr := us.conf.MasterAddresses
	page := 1
	recordPerPage := 100

	bytes := []byte(metaData)
	var amounts models.UserWalletAmounts

	if err := json.Unmarshal(bytes, &amounts); err != nil {
		return errors.New("Unmarshal error")
	}

	for {
		status, transactions := us.listTransactions(walletAddr, page, recordPerPage)
		if !status {
			log.Println("etherscan.io scan user-wallet return error")
			break
		}

		if len(transactions) == 0 {
			break
		}

		page = page + 1
		for _, transaction := range transactions {
			transactionObj := transaction.(map[string]interface{})
			contractAddress := transactionObj["contractAddress"].(string)
			txreceiptStatus := transactionObj["txreceipt_status"].(string)
			from := strings.ToLower(transactionObj["from"].(string))
			to := strings.ToLower(transactionObj["to"].(string))
			value := transactionObj["value"].(string)
			tnxTime, _ := strconv.ParseInt(transactionObj["timeStamp"].(string), 10, 64)

			if contractAddress != "" || txreceiptStatus != "1" || from != strings.ToLower(walletAddr) || value == "0" {
				continue
			}

			if tnxTime > expiredAt && tnxTime < startedAt {
				break
			}

			flagMasterAddr := false

			for _, masterAddr := range masterAddrArr {
				if to == strings.ToLower(masterAddr.Address) {
					flagMasterAddr = true
					continue
				}
			}

			if !flagMasterAddr {
				continue
			}

			for i := 0; i < len(amounts); i++ {
				if amounts[i].WeiValue == value {
					amounts[i].Status = "success"
				}
			}
			metaData, err := json.Marshal(amounts)
			if err != nil {
				return err
			}
			us.sendUserWalletHook(ID, walletAddr, to, string(metaData))
		}
	}
	return nil
}

func (us *UserService) listTransactions(walletAddr string, page int, offset int) (bool, []interface{}) {
	endpoint := us.conf.EtherscanURL
	apiKey := us.conf.EtherscanKey

	endpoint = fmt.Sprintf("%s?module=account&action=txlist&startblock=0&endblock=999999999&address=%s&page=%d&offset=%d&sort=desc&apikey=%s", endpoint, walletAddr, page, offset, apiKey)
	request, _ := http.NewRequest("GET", endpoint, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return false, nil
	}

	b, _ := ioutil.ReadAll(response.Body)

	var data map[string]interface{}
	json.Unmarshal(b, &data)

	status, ok := data["status"]
	message, _ := data["message"]
	result, _ := data["result"]

	if ok && status.(string) == "1" {
		return true, result.([]interface{})
	}
	log.Println(message)
	return false, nil
}

// ScanKYC : ...
func (us *UserService) ScanKYC() {
	users, err := us.ud.GetAllUsersNeedCheckKYC()
	if err == nil {
		for i := 0; i < len(*users); i++ {
			u := (*users)[i]
			status, errStr := us.checkPrimetrustContactID(u.PrimetrustContactID)
			if errStr != "404" {
				if status {
					us.sendKYCHook(u.ID, status, errStr)
				} else {
					if u.VerifiedLevel == 4 {
						us.sendKYCHook(u.ID, status, errStr)
					}
				}
			}
		}
	}
}

func (us *UserService) checkPrimetrustContactID(ID string) (bool, string) {
	endpoint := us.conf.PrimetrustEndpoint + "/contacts/" + ID
	request, _ := http.NewRequest("GET", endpoint, nil)
	request.Header.Add("Authorization", "Basic "+basicAuth(us.conf.PrimetrustUsername, us.conf.PrimetrustPassword))

	client := &http.Client{}
	response, err := client.Do(request)
	if err == nil {
		b, _ := ioutil.ReadAll(response.Body)
		var result map[string]interface{}
		json.Unmarshal([]byte(b), &result)

		if result["errors"] != nil {
			return false, "Primetrust id not found"
		}

		if result["data"] != nil {
			data := result["data"].(map[string]interface{})
			if data["attributes"] != nil {
				attributes := data["attributes"].(map[string]interface{})
				aml := attributes["aml-cleared"].(bool)
				cip := attributes["cip-cleared"].(bool)

				if aml && cip {
					return true, ""
				}

				if data["relationships"] != nil {
					relationships := data["relationships"].(map[string]interface{})
					if relationships["cip-checks"] != nil {
						cipChecks := relationships["cip-checks"].(map[string]interface{})
						if cipChecks["links"] != nil {
							links := cipChecks["links"].(map[string]interface{})

							related := links["related"].(string)
							if related != "" {
								arr := strings.Split(related, "/v2")
								end := us.conf.PrimetrustEndpoint + related
								if len(arr) == 2 {
									end = us.conf.PrimetrustEndpoint + arr[1]
								}
								req, _ := http.NewRequest("GET", end, nil)
								req.Header.Add("Authorization", "Basic "+basicAuth(us.conf.PrimetrustUsername, us.conf.PrimetrustPassword))
								client := &http.Client{}
								response, err := client.Do(req)
								if err == nil {
									b, _ := ioutil.ReadAll(response.Body)
									s := string(b[:])
									return false, s
								}

							}

						}
					}

				}
				return false, "404"
			}
		}

		return false, "Cannot parse data"
	}

	return false, "404"
}

func (us *UserService) sendKYCHook(userID uint, primetrustStatus bool, primetrustError string) error {
	jsonKYCData := make(map[string]interface{})
	jsonKYCData["PrimetrustContactStatus"] = primetrustStatus
	jsonKYCData["PrimetrustContactError"] = primetrustError
	jsonKYCData["ID"] = userID

	jsonData := make(map[string]interface{})
	jsonData["type"] = 4 /* WebhookTypeKYC */
	jsonData["data"] = jsonKYCData

	endpoint := us.conf.HookEndpoint
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
	} else {
		errStr := "Unknown"
		if hasMessage {
			errStr = message.(string)
		}
		return errors.New(errStr)
	}
}
