package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/constant-money/constant-event/config"
)

// HookService : struct
type HookService struct{}

// Event : send data to logic server
func (h HookService) Event(jsonData map[string]interface{}) error {
	conf := config.GetConfig()

	endpoint := conf.HookEndpoint
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
