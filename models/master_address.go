package models

import "github.com/jinzhu/gorm"

type MasterAddressStatus int

const (
	MasterAddressStatusDisable     MasterAddressStatus = 0
	MasterAddressStatusReady       MasterAddressStatus = 1
	MasterAddressStatusProgressing MasterAddressStatus = -1
	MasterAddressStatusWaiting     MasterAddressStatus = -2
)

type MasterAddress struct {
	gorm.Model

	Address         string `gorm:"not null;unique"`
	Status          MasterAddressStatus
	LastTnxHash     string
	LastTnxTime     int64
	LastBlockNumber int
	Nonce           int
}
