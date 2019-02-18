package daos

import (
	"github.com/constant-money/constant-event/models"
	"github.com/jinzhu/gorm"
)

// ReserveDAO : struct
type ReserveDAO struct {
	db *gorm.DB
}

// NewReserve : db
func NewReserve(database *gorm.DB) *ReserveDAO {
	return &ReserveDAO{db: database}
}

// FindAllReserves : ...
func (r *ReserveDAO) FindAllReserves() ([]*models.Reserve, error) {
	var (
		models []*models.Reserve
	)

	err := r.db.Table("reserves").Where("status in (?) and reserve_type in (?)", []int{1, 5}, []int{0, 1}).Find(&models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}
