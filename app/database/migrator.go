package database

import (
	"github.com/tryhd/dbo-test/app/types"
	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) {
	db.AutoMigrate(&types.Auth{})
	db.AutoMigrate(&types.Customer{})
	db.AutoMigrate(&types.Order{})
}
