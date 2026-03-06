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

type Database struct {
	db *gorm.DB
}

func NewDbase(database *gorm.DB) *Database {
	return &Database{
		db: database,
	}
}

func (repo *Database) GetID(id int) (*types.Model, error) {
	var person types.Model
	if err := repo.db.First(&person, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}
