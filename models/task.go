package models

import "github.com/jinzhu/gorm"

type TaskStatus int
type TaskMethod string
type NetworkID int

const (
	TaskStatusFailed      TaskStatus = 0
	TaskStatusSuccess     TaskStatus = 1
	TaskStatusPending     TaskStatus = -1
	TaskStatusRetry       TaskStatus = -2
	TaskStatusProgressing TaskStatus = -3
)

const (
	TaskMethodPurchase        TaskMethod = "purchase"
	TaskMethodRedeem          TaskMethod = "redeem"
	TaskMethodTransferByAdmin TaskMethod = "transfer_by_admin"
)

type Task struct {
	gorm.Model

	Method          TaskMethod
	Status          TaskStatus
	MasterAddress   string
	ContractAddress string
	ContractName    string
	Data            string
	Deleted         bool `gorm:"default:false"`
	NetworkID       uint
}
