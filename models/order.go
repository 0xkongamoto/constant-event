package models

import (
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
)

type MakerType int
type LocalStatus int

const (
	Buy MakerType = iota + 1
	Sell
)

const (
	StatusSellerInited LocalStatus = iota + 1
	StatusBuyerInited
	StatusWaitingBuyerSendMoney
	StatusBuyerSentMoney
	StatusSellerWaitBuyerSendMoney
	StatusDone
	StatusCancel
	StatusCancelByBot
)

// Order : struct
type Order struct {
	gorm.Model
	Type    MakerType
	User    wm.User `gorm:"foreignkey:UserID"`
	UserID  uint
	Shakers []Shaker `gorm:"foreignkey:OrderHistoryID;auto_preload:true"`
	Makers  []Maker  `gorm:"foreignkey:OrderHistoryID;auto_preload:true"`
}

// Maker : struct
type Maker struct {
	gorm.Model
	Type           MakerType
	UserID         uint
	Status         LocalStatus
	BkStatus       LocalStatus
	Shakers        []Shaker `gorm:"foreignkey:MakerID;auto_preload:true"`
	OrderHistoryID uint
}

// Shaker : struct
type Shaker struct {
	gorm.Model
	Type           MakerType
	Status         LocalStatus
	BkStatus       LocalStatus
	UserID         uint
	MakerID        uint
	Maker          Maker
	OrderHistoryID uint
	DealTime       int64
}
