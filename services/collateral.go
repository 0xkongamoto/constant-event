package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	bedaos "github.com/constant-money/constant-web/backend/daos"
	"github.com/jinzhu/gorm"
)

type CollateralService struct {
	cd          *bedaos.Collateral
	db          *gorm.DB
	RateFeeding bool
}

func NewCollateralService(db *gorm.DB, cd *bedaos.Collateral) *CollateralService {
	return &CollateralService{
		cd:          cd,
		db:          db,
		RateFeeding: false,
	}
}

func (cs *CollateralService) RateFeed() {
	collaterals, err := cs.cd.FindAll()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, c := range collaterals {
		fmt.Printf("Feeding %s\n", c.Symbol)

		req, err := http.NewRequest("GET", fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=USD", c.Symbol), nil)

		if err != nil {
			fmt.Println("Init request failed", c.Symbol, err.Error())
			continue
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Call external service failed", c.Symbol, err.Error())
			continue
		}
		defer func(r *http.Response) {
			err := r.Body.Close()
			if err != nil {
				fmt.Println("Close body failed", c.Symbol, err.Error())
			}
		}(res)

		if res.StatusCode != http.StatusOK {
			fmt.Println("Response status != 200")
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Read body failed", c.Symbol, err.Error())
			continue
		}

		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Unmarshal body failed", c.Symbol, err.Error())
			continue
		}

		_, exist := data[c.Symbol]
		if exist {
			val := data[c.Symbol].(float64)
			intVal := uint64(val * 100)
			// todo call update symbol
			if intVal != c.Value {
				c.Value = intVal
				err := cs.cd.Update(cs.db, c)
				fmt.Printf("Update rate for symbol %s failed %s\n", c.Symbol, err.Error())
			}
		}

		time.Sleep(1 * time.Second)
	}
}
