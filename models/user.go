package models

import "github.com/jinzhu/gorm"

// UserWalletStatus :
type UserWalletStatus int

// UserWalletStatus :
const (
	UserWalletStatusPending UserWalletStatus = iota
	UserWalletStatusDone
)

// UserWallet :
type UserWallet struct {
	gorm.Model

	UserID uint

	WalletAddress string
	Metadata      string
	Active        int `gorm:"default:0"`
	Status        UserWalletStatus
	ExpiredAt     int64
	StartedAt     int64
}

// UserWalletAmount :
type UserWalletAmount struct {
	WeiValue string
	Status   string
}

// UserWalletAmounts :
type UserWalletAmounts []UserWalletAmount

// User : struct
type User struct {
	gorm.Model

	VerifiedLevel           int
	PrimetrustContactID     string
	PrimetrustContactError  string `gorm:"type:text"`
	PrimetrustContactStatus int    `gorm:"default:-1"`
}
