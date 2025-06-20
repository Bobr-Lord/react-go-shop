package service

import "github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"

func (s *Service) AddCartItem(idItem string, idUser string) error {
	return s.r.AddCartItem(idItem, idUser)
}

func (s *Service) DeleteCartItem(idItem string, idUser string) error {
	return s.r.DeleteCartItem(idItem, idUser)
}

func (s *Service) GetCartItems(idUser string) ([]models.ProductWithQuantity, error) {
	return s.r.GetCartItems(idUser)
}
