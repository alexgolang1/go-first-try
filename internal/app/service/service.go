package service

import (
	"back-api/internal/app/types"
)

type Repository interface {
	GetIDSrv(id int) (*types.Model, error)
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}
