package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CollateralLoanStatus int
type CollateralLoanInterestType int
type CollateralWalletType int
type CollateralLoanTransactionType int

const (
	CollateralLoanStatusPending = iota
	CollateralLoanStatusAccepted
	CollateralLoanStatusPayingInterest
	CollateralLoanStatusDone
	CollateralLoanStatusCancelled
	CollateralLoanStatusWithdrawed
)

type CollateralLoan struct {
	gorm.Model

	UserID uint
	User   *User

	Amount               uint64
	CollateralAddress    string
	CollateralPrivateKey string
	Status               CollateralLoanStatus
	CollateralID         uint
	Collateral           *Collateral
	CollateralAmount     uint64
	CollateralValue      uint64
	CollateralThreshold  uint
	CollateralReturn     uint
	InterestRate         uint
	InterestAmount       uint64
	InterestType         CollateralLoanInterestType
	InterestAuto         uint
	Period               uint
	DelayDuration        uint
	WithdrawAddress      string
	WithdrawAt           *time.Time
	StartAt              *time.Time
	NextPayAt            *time.Time
	FinishAt             *time.Time
	CancelReason         string `gorm:"type:text"`

	CollateralLoanTransactions []*CollateralLoanTransaction
}

type Collateral struct {
	gorm.Model

	Name       string
	Symbol     string
	Address    string
	Decimals   uint
	WalletType CollateralWalletType

	Value uint64
}

type CollateralLoanTransaction struct {
	gorm.Model

	CollateralLoanID uint
	CollateralLoan   *CollateralLoan
	Type             CollateralLoanTransactionType
	Amount           uint64
	Description      string
}
