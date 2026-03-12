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

type DataBase struct {
	db *gorm.DB
}

func NewDbase(database *gorm.DB) *DataBase {
	return &DataBase{
		db: database,
	}
}

func (repo *DataBase) GetUser(id int) (*types.Model, error) {
	var person types.Model
	if err := repo.db.First(&person, id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

func (repo *DataBase) CreateUser(name, surname string) error {
	new_person := types.Model{
		Name:    name,
		Surname: surname,
	}

	return repo.db.Create(&new_person).Error
}

func (repo *DataBase) DeleteUser(id int) error {
	return repo.db.Delete(&types.Model{}, id).Error
}

func (repo *DataBase) UpdateUser(id int, name, surname string) error {
	return repo.db.Model(&types.Model{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":    name,
		"surname": surname,
	}).Error
}
