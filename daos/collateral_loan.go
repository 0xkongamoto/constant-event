package daos

import (
	"github.com/constant-money/constant-event/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type CollateralLoanDAO struct {
	db *gorm.DB
}

// InitCollateralLoanDAO :
func InitCollateralLoanDAO(database *gorm.DB) *CollateralLoanDAO {
	return &CollateralLoanDAO{
		db: database,
	}
}

func (cl *CollateralLoanDAO) FindAllPending(limit int) ([]*models.CollateralLoan, error) {
	var (
		collateralLoans []*models.CollateralLoan
	)

	query := cl.db.Table("collateral_loans").
		Where("status = ?", models.CollateralLoanStatusPending).
		Order("id desc").
		Limit(limit)

	if err := query.Find(&collateralLoans).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}
	return collateralLoans, nil
}

func (cl *CollateralLoanDAO) Update(tx *gorm.DB, model *models.CollateralLoan) error {
	if err := tx.Save(model).Error; err != nil {
		return errors.Wrap(err, "tx.Update.CollateralLoan")
	}
	return nil
}
