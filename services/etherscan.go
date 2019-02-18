package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/constant-money/constant-web/event/config"
)

// EtherscanService :
type EtherscanService struct{}

// ListTransactions :
func (s EtherscanService) ListTransactions(address string, page int, offset int) (bool, []interface{}) {
	conf := config.GetConfig()
	endpoint := conf.EtherscanURL
	apiKey := conf.EtherscanKey

	endpoint = fmt.Sprintf("%s?module=account&action=txlist&startblock=0&endblock=999999999&address=%s&page=%d&offset=%d&sort=asc&apikey=%s", endpoint, address, page, offset, apiKey)

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
