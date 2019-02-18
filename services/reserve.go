package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/constant-money/constant-web/event/config"
	"github.com/constant-money/constant-web/event/daos"
	"github.com/constant-money/constant-web/event/models"
)

// ReserveService : struct
type ReserveService struct {
	rd      *daos.ReserveDAO
	conf    *config.Config
	Running bool
}

// InitReserveService : reserveDAO
func InitReserveService(reserveDAO *daos.ReserveDAO, config *config.Config) *ReserveService {
	return &ReserveService{
		rd:   reserveDAO,
		conf: config,
	}
}

// PrimetrustHook : ...
func (r *ReserveService) PrimetrustHook() {
	result, err := r.rd.FindAllReserves()
	if err == nil {
		for i := 0; i < len(result); i++ {
			reserve := result[i]
			go r.execLogic(reserve)
		}
	}
}

func (r *ReserveService) execLogic(reserve *models.Reserve) {
	var cons string
	if reserve.ReserveType == 0 {
		tmp, err := r.contributions(reserve.ExtID)
		if err != nil {
			return
		}
		cons = tmp
	} else {
		tmp, err := r.disbursement(reserve.ExtID)
		if err != nil {
			return
		}
		cons = tmp
	}

	status, err := r.fundTransfer(cons)
	if err != nil {
		return
	}

	if status == "settled" {
		jsonData := make(map[string]interface{})
		jsonData["Action"] = status
		jsonData["ID"] = reserve.ID

		jsonWebhook := make(map[string]interface{})
		jsonWebhook["type"] = 1 /* WebhookTypeReserve */
		jsonWebhook["data"] = jsonData

		_, err := r.hook(jsonWebhook)
		if err != nil {
			fmt.Println("Cannot call hook data")
		}
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (r *ReserveService) contributions(extID string) (string, error) {
	endpoint := r.conf.PrimetrustEndpoint + "/contributions/" + extID
	request, _ := http.NewRequest("GET", endpoint, nil)
	request.Header.Add("Authorization", "Basic "+basicAuth(r.conf.PrimetrustUsername, r.conf.PrimetrustPassword))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	b, _ := ioutil.ReadAll(response.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)

	if result["errors"] != nil {
		return "", errors.New("ExtID not found")
	}

	if result["data"] != nil {
		data := result["data"].(map[string]interface{})
		if data["relationships"] != nil {
			relationships := data["relationships"].(map[string]interface{})
			if relationships["funds-transfer"] != nil {
				funds := relationships["funds-transfer"].(map[string]interface{})
				if funds["links"] != nil {
					links := funds["links"].(map[string]interface{})
					related := links["related"].(string)
					if related != "" {
						s := strings.Split(related, "/")
						if len(s) > 0 {
							final := s[len(s)-1]
							return final, nil
						}
					}
				}
			}
		}
	}

	return "", errors.New("Cannot parse data")
}

func (r *ReserveService) disbursement(extID string) (string, error) {
	endpoint := r.conf.PrimetrustEndpoint + "/disbursements/" + extID
	request, _ := http.NewRequest("GET", endpoint, nil)
	request.Header.Add("Authorization", "Basic "+basicAuth(r.conf.PrimetrustUsername, r.conf.PrimetrustPassword))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	b, _ := ioutil.ReadAll(response.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)

	if result["errors"] != nil {
		return "", errors.New("ExtID not found")
	}

	if result["data"] != nil {
		data := result["data"].(map[string]interface{})
		if data["relationships"] != nil {
			relationships := data["relationships"].(map[string]interface{})
			if relationships["funds-transfer"] != nil {
				funds := relationships["funds-transfer"].(map[string]interface{})
				if funds["links"] != nil {
					links := funds["links"].(map[string]interface{})
					related := links["related"].(string)
					if related != "" {
						s := strings.Split(related, "/")
						if len(s) > 0 {
							final := s[len(s)-1]
							return final, nil
						}
					}
				}
			}
		}
	}

	return "", errors.New("Cannot parse data")
}

func (r *ReserveService) fundTransfer(cons string) (string, error) {
	endpoint := r.conf.PrimetrustEndpoint + "/funds-transfers/" + cons
	request, _ := http.NewRequest("GET", endpoint, nil)
	request.Header.Add("Authorization", "Basic "+basicAuth(r.conf.PrimetrustUsername, r.conf.PrimetrustPassword))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	b, _ := ioutil.ReadAll(response.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)

	if result["errors"] != nil {
		return "", errors.New("Cons not found")
	}

	if result["data"] != nil {
		data := result["data"].(map[string]interface{})
		if data["attributes"] != nil {
			attributes := data["attributes"].(map[string]interface{})
			status := attributes["status"].(string)
			return status, nil
		}
	}

	return "", errors.New("Cannot parse data")
}

func (r *ReserveService) hook(jsonData map[string]interface{}) (string, error) {
	endpoint := r.conf.HookEndpoint
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	b, _ := ioutil.ReadAll(response.Body)

	var data map[string]interface{}
	json.Unmarshal(b, &data)

	status, ok := data["status"]
	message, hasMessage := data["message"]

	if ok && status.(float64) > 0 {
		return "", nil
	} else {
		errStr := "Unknown"
		if hasMessage {
			errStr = message.(string)
		}
		return "", errors.New(errStr)
	}
}
