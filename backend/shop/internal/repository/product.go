package repository

import (
	"fmt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"
)

func (r *Repository) CreateProduct(req *models.CreateProductRequest) error {
	query := fmt.Sprintf("INSERT INTO %v ( name, description, price, category, image_url) VALUES ($1, $2, $3, $4, $5)", productTableName)
	_, err := r.db.Exec(query, req.Name, req.Description, req.Price, req.Category, req.ImageUrl)
	return err
}

func (r *Repository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	query := fmt.Sprintf("SELECT id, name, description, price, category, image_url FROM %v", productTableName)
	err := r.db.Select(&products, query)
	fmt.Println(products)
	return products, err
}
