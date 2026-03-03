package service

import (
	"back-api/internal/app/types"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func New(database *gorm.DB) *Service {
	return &Service{
		db: database,
	}
}

func (srv *Service) GetIDSrv(id int) (*types.Model, error) {
	var person types.Model

	if err := srv.db.First(&person, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}
