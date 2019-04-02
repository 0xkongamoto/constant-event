package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/constant-money/constant-web-api/serializers"
	"github.com/constant-money/constant-web-api/services/3rd/primetrust"
)

// UserService : struct
type UserService struct {
	primetrust   *primetrust.Primetrust
	hookEndpoint string
	ptEndpoint   string
}

// InitUserService :
func InitUserService(primetrust *primetrust.Primetrust, hookEndpoint string, ptEndpoint string) *UserService {
	return &UserService{
		primetrust:   primetrust,
		hookEndpoint: hookEndpoint,
		ptEndpoint:   ptEndpoint,
	}
}

// CheckPrimetrustContactID : contactID
func (us *UserService) CheckPrimetrustContactID(ID string) (bool, string) {
	response, err := us.primetrust.GetContactByID(ID)

	if err != nil {
		return false, "404"
	}

	contactData := response.Data
	if contactData != nil {
		attributes := contactData.Attributes
		if attributes != nil {
			aml := attributes.AMLCleared
			cip := attributes.CIPCleared

			if aml && cip {
				return true, ""
			}

			relationships := contactData.Relationships
			if relationships != nil {
				cipChecks := relationships.CIPChecks
				if cipChecks != nil {
					links := cipChecks.Links
					if links != nil {
						related := links.Related
						if related != "" {
							arr := strings.Split(related, "/v2")
							end := us.ptEndpoint + related
							if len(arr) == 2 {
								end = us.ptEndpoint + arr[1]
							}
							req, _ := http.NewRequest("GET", end, nil)
							token, err := us.primetrust.GetToken()
							if err != nil {
								return false, "404"
							}

							req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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

// SendKYCHook : userID, primetrustStatus, primetrustError
func (us *UserService) SendKYCHook(userID uint, primetrustStatus bool, primetrustError string) error {
	jsonKYCData := make(map[string]interface{})
	jsonKYCData["PrimetrustContactStatus"] = primetrustStatus
	jsonKYCData["PrimetrustContactError"] = primetrustError
	jsonKYCData["ID"] = userID

	jsonData := make(map[string]interface{})
	jsonData["type"] = serializers.WebhookTypeKYC /* WebhookTypeKYC */
	jsonData["data"] = jsonKYCData

	endpoint := us.hookEndpoint
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

	e := data["Error"]

	if e == nil {
		return nil
	}

	return errors.New(e.(string))
}
