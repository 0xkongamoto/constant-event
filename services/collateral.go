package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
	"github.com/jinzhu/gorm"
)

type CollateralService struct {
	cd          *daos.CollateralDAO
	db          *gorm.DB
	RateFeeding bool
}

func NewCollateralService(db *gorm.DB, cd *daos.CollateralDAO) *CollateralService {
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

		_, exist := data["USD"]
		if exist {
			val := data["USD"].(float64)
			intVal := uint64(val * 100)
			// todo call update symbol
			if intVal != c.Value {
				errTx := models.WithTransaction(func(tx *gorm.DB) error {
					c.Value = intVal
					if err := cs.cd.Update(tx, c); err != nil {
						log.Printf("Update rate for symbol %s failed %s\n", c.Symbol, err.Error())
						return err
					}
					return nil
				})

				if errTx != nil {
					log.Printf("Update Tnx rate for symbol %s failed %s\n", c.Symbol, errTx.Error())
				}
			}
		} else {
			log.Printf("Not found symbol %s \n", c.Symbol)
		}
		time.Sleep(1 * time.Second)
	}
}
