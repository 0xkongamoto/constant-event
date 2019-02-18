package daos

import (
	"github.com/constant-money/constant-event/models"
	"github.com/jinzhu/gorm"
)

// OrderDAO : struct
type OrderDAO struct {
	db *gorm.DB
}

// InitOrderDAO : ...
func InitOrderDAO(database *gorm.DB) *OrderDAO {
	return &OrderDAO{
		db: database,
	}
}

// GetAllOrders : ...
func (od *OrderDAO) GetAllOrders() ([]*models.Order, error) {
	orders := []*models.Order{}
	err := models.Database().Preload("User").Preload("Makers").Preload("Shakers").Preload("Shakers.Maker").
		Joins("LEFT JOIN shakers ON shakers.order_history_id=orders.id").
		Joins("LEFT JOIN makers ON makers.order_history_id=orders.id").
		Where("shakers.status in (?)", []models.LocalStatus{models.StatusWaitingBuyerSendMoney, models.StatusSellerWaitBuyerSendMoney}).Group("orders.ID").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}
