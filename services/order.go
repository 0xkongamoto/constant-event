package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/constant-money/constant-web/event/config"
	"github.com/constant-money/constant-web/event/daos"
)

const (
	WindowTime = 1800 // 30 minutes
)

// OrderService : struct
type OrderService struct {
	od      *daos.OrderDAO
	conf    *config.Config
	Running bool
}

// InitOrderService : orderDAO
func InitOrderService(orderDao *daos.OrderDAO, cf *config.Config) *OrderService {
	return &OrderService{
		od:   orderDao,
		conf: cf,
	}
}

// ScanOrders : ...
func (o *OrderService) ScanOrders() {
	orders, _ := o.od.GetAllOrders()
	for i := 0; i < len(orders); i++ {
		order := orders[i]
		shakers := order.Shakers
		for j := 0; j < len(shakers); j++ {
			shaker := shakers[j]
			now := time.Now().UnixNano() / int64(time.Millisecond)
			if (now-shaker.DealTime)/1000 >= WindowTime {
				fmt.Println("Fire hook", order.ID, shaker.OrderHistoryID)
				go o.sendHook(order.ID, shaker.OrderHistoryID)
			}
		}
	}
}

func (o *OrderService) sendHook(orderID uint, canceledOrderID uint) error {
	jsonData := make(map[string]interface{})
	jsonData["ID"] = orderID
	jsonData["CanceledOrderID"] = canceledOrderID

	jsonWebhook := make(map[string]interface{})
	jsonWebhook["type"] = 2 /* WebhookTypeOrder */
	jsonWebhook["data"] = jsonData

	endpoint := o.conf.HookEndpoint
	endpoint = fmt.Sprintf("%s", endpoint)
	jsonValue, _ := json.Marshal(jsonWebhook)

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
