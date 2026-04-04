package types

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Password string
	Email    string
}

type LoginRequset struct {
	Password string
	Email    string
}

type LoginResponse struct {
	Token string
}
