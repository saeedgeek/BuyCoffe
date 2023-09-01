package db

import (
	Base "BuyCoffee/cmd"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(Base.SqliteDsn), &gorm.Config{})
	if err != nil {
		panic("Cannot open sqlite, Error: " + err.Error())
	}
}
