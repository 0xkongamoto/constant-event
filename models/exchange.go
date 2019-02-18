package models

import (
	"encoding/xml"

	"github.com/jinzhu/gorm"
)

// Exchange : struct
type Exchange struct {
	gorm.Model
	CurrencyCode string
	CurrencyName string
	Buy          uint64
	Transfer     uint64
	Sell         uint64
}

type ExchangeRes struct {
	XMLName       xml.Name       `xml:"ExrateList"`
	ExchangeRates []ExchangeRate `xml:"Exrate"`
}

type ExchangeRate struct {
	XMLName      xml.Name `xml:"Exrate"`
	CurrencyCode string   `xml:"CurrencyCode,attr"`
	CurrencyName string   `xml:"CurrencyName,attr"`
	Buy          string   `xml:"Buy,attr"`
	Transfer     string   `xml:"Transfer,attr"`
	Sell         string   `xml:"Sell,attr"`
}

// TableName : exchange
func (e Exchange) TableName() string {
	return "exchange_rate"
}
