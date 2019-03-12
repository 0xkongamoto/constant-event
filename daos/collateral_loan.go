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

	query := cl.db.Table("collateral_loans").Preload("Collateral").
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
	// HOUR(next_pay_at) = HOUR(now())
	query := cl.db.Raw(`SELECT *
						FROM collateral_loans 
						WHERE 
							status = ? AND 
							YEAR(next_pay_at) = YEAR(now()  + interval ? day) AND 
							MONTH(next_pay_at) = MONTH(now()  + interval ? day) AND 
							DAY(next_pay_at) = DAY(now() + interval ? day)
						LIMIT ? 
						OFFSET ?`, wm.CollateralLoanStatusPayingInterest, dayNumber, dayNumber, dayNumber, limit, offset)

	if err := query.Scan(&collateralLoans).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}
	return collateralLoans, nil
}

func (cl *CollateralLoanDAO) FindAllPayingInterestByDay(dayNumber int, page int, limit int) ([]*wm.CollateralLoan, error) {
	var (
		collateralLoans []*wm.CollateralLoan
		offset          = page*limit - limit
	)

	query := cl.db.Raw(`SELECT *
						FROM collateral_loans 
						WHERE 
							status = ? AND 
							YEAR(next_pay_at) <= YEAR(now() - interval ? day) AND 
							MONTH(next_pay_at) <= MONTH(now() - interval ? day) AND 
							DAY(next_pay_at) <= DAY(now() - interval ? day)
						LIMIT ? 
						OFFSET ?`, wm.CollateralLoanStatusAccepted, dayNumber, dayNumber, dayNumber, limit, offset)

	if err := query.Scan(&collateralLoans).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}

	return collateralLoans, nil
}

func (cl *CollateralLoanDAO) FindAllDowntrend(amount uint64, page int, limit int) ([]*wm.CollateralLoan, error) {
	var (
		collateralLoans []*wm.CollateralLoan
		offset          = page*limit - limit
	)

	query := cl.db.Table("collateral_loans").Preload("Collateral").Preload("CollateralLoanTransactions").
		Joins("JOIN collateral_loan_transactions ON collateral_loan_transactions.collateral_loan_id=collateral_loans.id").
		Where(`	(status = ? OR status = ?) AND 
				collateral_loan_transactions.collateral_rate >= ? AND
				collateral_loan_transactions.type = ?
			`, wm.CollateralLoanStatusPayingInterest, wm.CollateralLoanStatusAccepted, amount, wm.CollateralLoanTransactionTypeReceive).
		Order("id desc").
		Limit(limit).
		Offset(offset)

	if err := query.Find(&collateralLoans).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}
	return collateralLoans, nil
}
