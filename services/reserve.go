package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
	"github.com/constant-money/constant-web-api/serializers"
	"github.com/constant-money/constant-web-api/services/3rd/primetrust"
	"github.com/pkg/errors"
)

// ReserveService : struct
type ReserveService struct {
	primetrust   *primetrust.Primetrust
	rd           *daos.ReserveDAO
	hookEndpoint string
	Running      bool
}

// InitReserveService : reserveDAO
func InitReserveService(reserveDAO *daos.ReserveDAO, primetrust *primetrust.Primetrust, hookEndpoint string) *ReserveService {
	return &ReserveService{
		rd:           reserveDAO,
		primetrust:   primetrust,
		hookEndpoint: hookEndpoint,
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
		jsonWebhook["type"] = serializers.WebhookTypeReserve
		jsonWebhook["data"] = jsonData

		_, err := r.hook(jsonWebhook)
		if err != nil {
			fmt.Println("Cannot call hook data")
		}
	}
}

func (r *ReserveService) contributions(extID string) (string, error) {
	response, err := r.primetrust.GetContribution(extID)
	if err != nil {
		return "", errors.New("Cannot parse data")
	}

	contributionData := response.Data
	if contributionData != nil {
		relationships := contributionData.Relationships
		if relationships != nil {
			fundTransfer := relationships.FundsTransfer
			if fundTransfer != nil {
				links := fundTransfer.Links
				if links != nil {
					related := links.Related
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
	response, err := r.primetrust.GetDisbursement(extID)
	if err != nil {
		return "", errors.New("Cannot parse data")
	}

	disbursementData := response.Data
	if disbursementData != nil {
		relationships := disbursementData.Relationships
		if relationships != nil {
			fundTransfer := relationships.FundsTransfer
			if fundTransfer != nil {
				links := fundTransfer.Links
				if links != nil {
					related := links.Related
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
	response, err := r.primetrust.GetFundsTransfer(cons)
	if err != nil {
		return "", errors.New("Cannot parse data")
	}

	fundTransferData := response.Data
	if fundTransferData != nil {
		attributes := fundTransferData.Attributes
		if attributes != nil {
			status := attributes.Status
			return status, nil
		}
	}

	return "", errors.New("Cannot parse data")
}

func (r *ReserveService) hook(jsonData map[string]interface{}) (string, error) {
	endpoint := r.hookEndpoint
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
