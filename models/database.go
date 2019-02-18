package models

import (
	"log"

	"github.com/constant-money/constant-web/event/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbInst *gorm.DB

// Database : gorm.DB
func Database() *gorm.DB {
	if dbInst == nil {
		conf := config.GetConfig()
		d, err := gorm.Open("mysql", conf.Db)

		d.LogMode(false)

		if err != nil {
			log.Println(err)
			return nil
		}

		dbInst = d.Set("gorm.save_associations", false)
		dbInst.DB().SetMaxOpenConns(20)
		dbInst.DB().SetMaxIdleConns(10)
	}
	return dbInst
}

func WithTransaction(callback func(*gorm.DB) error) error {
	tx := dbInst.Begin()
	if err := callback(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func WithDB(callback func(*gorm.DB) error) error {
	if err := callback(dbInst); err != nil {
		return err
	}

	return nil
}
