package types

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	Name    string
	Surname string
}
