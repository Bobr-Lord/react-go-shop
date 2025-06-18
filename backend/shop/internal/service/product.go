package service

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/response"
	"net/http"
)

func (s *Service) CreateProduct(req *models.CreateProductRequest) error {
	err := s.r.CreateProduct(req)
	if err != nil {
		return response.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (s *Service) GetAllProducts() ([]models.Product, error) {
	products, err := s.r.GetAllProducts()
	if err != nil {
		return nil, response.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return products, nil
}
