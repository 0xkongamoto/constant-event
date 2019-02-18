package models

import "github.com/jinzhu/gorm"

// Reserve : struct
type Reserve struct {
	gorm.Model

	Status      int
	ReserveType int
	ExtID       string
}
