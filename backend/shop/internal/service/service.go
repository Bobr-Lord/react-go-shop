package service

import "github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/repository"

type Service struct {
	r *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{r}
}
