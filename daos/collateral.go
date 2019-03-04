package daos

import (
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type CollateralDAO struct {
	db *gorm.DB
}

// InitCollateralDAO :
func InitCollateralDAO(database *gorm.DB) *CollateralDAO {
	return &CollateralDAO{
		db: database,
	}
}

func (r *CollateralDAO) Create(tx *gorm.DB, model *wm.Collateral) error {
	if err := tx.Create(model).Error; err != nil {
		return errors.Wrap(err, "tx.Create")
	}
	return nil
}

func (r *CollateralDAO) Update(tx *gorm.DB, model *wm.Collateral) error {
	if err := tx.Save(model).Error; err != nil {
		return errors.Wrap(err, "tx.Update")
	}
	return nil
}

func (r *CollateralDAO) Delete(tx *gorm.DB, model *wm.Collateral) error {
	if err := tx.Delete(model).Error; err != nil {
		return errors.Wrap(err, "tx.Delete")
	}
	return nil
}

func (r *CollateralDAO) FindByID(id uint) (*wm.Collateral, error) {
	var model wm.Collateral
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, errors.Wrap(err, "db.First")
	}
	return &model, nil
}

func (r *CollateralDAO) FindAll() ([]*wm.Collateral, error) {
	var (
		models []*wm.Collateral
		//offset = page*limit - limit
	)

	query := r.db.Table("collaterals").Order("id desc")
	//query = query.Limit(limit).Offset(offset)
	//if filter != nil {
	//	query = query.Where(*filter)
	//}

	if err := query.Find(&models).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}
	return models, nil
}
