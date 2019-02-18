package daos

import (
	"github.com/constant-money/constant-web/event/models"
	"github.com/jinzhu/gorm"
)

// ExchangeDAO : struct
type ExchangeDAO struct {
	db *gorm.DB
}

// InitExchangeDAO : ...
func InitExchangeDAO(database *gorm.DB) *ExchangeDAO {
	return &ExchangeDAO{
		db: database,
	}
}

// Create : exchange
func (e *ExchangeDAO) Create(ex *models.Exchange) error {
	return e.db.Create(ex).Error
}

// Update : exchange
func (e *ExchangeDAO) Update(ex *models.Exchange) error {
	return e.db.Save(ex).Error
}

// ExchangeRate :  vietcombank
func (e *ExchangeDAO) ExchangeRate() *models.Exchange {
	exchange := models.Exchange{}
	e.db.First(&exchange)
	return &exchange
}
