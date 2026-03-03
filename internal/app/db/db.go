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

type Repository struct {
	db *gorm.DB
}

func NewRepo(database *gorm.DB) *Repository {
	return &Repository{
		db: database,
	}
}

func (r *Repository) GetIDSrv(id int) (*types.Model, error) {
	var person types.Model
	if err := r.db.First(&person, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}
