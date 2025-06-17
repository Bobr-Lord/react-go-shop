package repository

import "github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"

func (r *Repository) CreateProduct(req *models.CreateProductRequest) error {
	query := `INSERT INTO` + productTableName + ` ( name, description, price, category, image_url) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, req.Name, req.Description, req.Price, req.Category, req.ImageUrl)
	return err
}
