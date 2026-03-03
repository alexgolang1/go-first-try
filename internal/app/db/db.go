package db

import (
	"back-api/internal/app/types"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dsn := "any param"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&types.Model{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
