package services

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
)

// ExchangeService : struct
type ExchangeService struct {
	ed      *daos.ExchangeDAO
	Running bool
}

// InitExchangeService : exchangeDAO
func InitExchangeService(exchangeDAO *daos.ExchangeDAO) *ExchangeService {
	return &ExchangeService{
		ed: exchangeDAO,
	}
}

// ParseExchangeRate : ...
func (e *ExchangeService) ParseExchangeRate() {
	request, _ := http.NewRequest("GET", "https://www.vietcombank.com.vn/ExchangeRates/ExrateXML.aspx", nil)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	b, _ := ioutil.ReadAll(response.Body)
	var exchange models.ExchangeRes
	xml.Unmarshal(b, &exchange)

	rates := exchange.ExchangeRates
	ex := e.ed.ExchangeRate()
	for i := 0; i < len(rates); i++ {
		rate := rates[i]
		if rate.CurrencyCode == "USD" {
			ex.CurrencyCode = rate.CurrencyCode
			ex.CurrencyName = rate.CurrencyName
			v, _ := strconv.ParseInt(rate.Buy, 10, 64)
			ex.Buy = uint64(v)
			v, _ = strconv.ParseInt(rate.Transfer, 10, 64)
			ex.Transfer = uint64(v)
			v, _ = strconv.ParseInt(rate.Sell, 10, 64)
			ex.Sell = uint64(v)

			if ex.ID > 0 {
				e.ed.Update(ex)
			} else {
				e.ed.Create(ex)
			}

			break
		}
	}
}
