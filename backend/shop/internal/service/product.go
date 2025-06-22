package service

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/response"
	"net/http"
	"slices"
)

func (s *Service) CreateProduct(req *models.CreateProductRequest) (string, error) {
	id, err := s.r.CreateProduct(req)
	if err != nil {
		return "", response.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return id, nil
}

func (s *Service) GetAllProducts() ([]models.Product, error) {
	products, err := s.r.GetAllProducts()
	if err != nil {
		return nil, response.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	slices.Reverse(products)
	return products, nil
}

func (s *Service) DeleteProduct(id string) error {
	err := s.r.DeleteProduct(id)
	if err != nil {
		return response.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (s *Service) GetAllProductsPrivate(idUser string) ([]models.Product, error) {
	return s.r.GetAllProductsPrivate(idUser)
}

func (s *Service) DecrementProduct(idUser string, idProduct string) error {
	return s.r.DecrementProduct(idUser, idProduct)
}
