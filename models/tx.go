package models

import "github.com/jinzhu/gorm"

// Tx :
type Tx struct {
	gorm.Model

	Hash string

	Payload   string
	Status    int
	ChainID   uint
	NetworkID uint
	Offchain  string
	TaskID    uint

	MasterAddress   string
	ContractAddress string
	ContractMethod  string

	Address string
}
