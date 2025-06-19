package repository

import (
	"fmt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"
)

func (r *Repository) CreateProduct(req *models.CreateProductRequest) (string, error) {
	query := fmt.Sprintf("INSERT INTO %v ( name, description, price, category, image_url) VALUES ($1, $2, $3, $4, $5) RETURNING id", productTableName)
	var id string
	err := r.db.QueryRow(query, req.Name, req.Description, req.Price, req.Category, req.ImageUrl).Scan(&id)
	return id, err
}

func (r *Repository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	query := fmt.Sprintf("SELECT id, name, description, price, category, image_url FROM %v", productTableName)
	err := r.db.Select(&products, query)
	return products, err
}

func (r *Repository) DeleteProduct(id string) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id = $1", productTableName)
	_, err := r.db.Exec(query, id)
	return err
}
