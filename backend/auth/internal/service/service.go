package service

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/config"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/repository"
)

type Service struct {
	repo *repository.Repository
	cfg  *config.Config
}

func NewService(repo *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}
