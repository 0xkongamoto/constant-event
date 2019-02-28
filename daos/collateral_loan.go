package daos

import (
	wm "github.com/constant-money/constant-web-api/models"
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

func (cl *CollateralLoanDAO) Update(tx *gorm.DB, model *wm.CollateralLoan) error {
	if err := tx.Save(model).Error; err != nil {
		return errors.Wrap(err, "tx.Update.CollateralLoan")
	}
	return nil
}

func (cl *CollateralLoanDAO) FindAllPending(lastIndex uint, limit int) ([]*wm.CollateralLoan, error) {
	var (
		collateralLoans []*wm.CollateralLoan
	)

	query := cl.db.Table("collateral_loans").
		Where("status = ? AND id > ?", wm.CollateralLoanStatusPending, lastIndex).
		Order("id desc").
		Limit(limit)

	if err := query.Find(&collateralLoans).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}
	return collateralLoans, nil
}

func (cl *CollateralLoanDAO) FindAllPayingByDate(dayNumber uint, page int, limit int) ([]*wm.CollateralLoan, error) {
	var (
		collateralLoans []*wm.CollateralLoan
		offset          = page*limit - limit
	)

	query := cl.db.Raw(`SELECT *
						FROM collateral_loans 
						WHERE 
							status = ? AND 
							MONTH(next_pay_at) = MONTH(now()  + interval ? day) AND 
							YEAR(next_pay_at) = YEAR(now()  + interval ? day) AND 
							DAY(next_pay_at) = DAY(now() + interval ? day) AND
							DAY(next_pay_at) = DAY(now() + interval ? day) AND
							HOUR(next_pay_at) = HOUR(now()) 
						LIMIT ? 
						OFFSET ?`, wm.CollateralLoanStatusPayingInterest, dayNumber, dayNumber, dayNumber, dayNumber, limit, offset)

	if err := query.Scan(&collateralLoans).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}
	return collateralLoans, nil
}
