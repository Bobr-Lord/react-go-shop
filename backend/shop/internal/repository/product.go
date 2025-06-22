package repository

import (
	"fmt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/response"
	"net/http"
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

func (r *Repository) GetAllProductsPrivate(idUser string) ([]models.Product, error) {
	query := `
		SELECT 
			p.id,
			p.name,
			p.description,
			p.price,
			p.category,
			p.image_url,
			COALESCE(ci.quantity, 0) AS quantity
		FROM products p
		LEFT JOIN cart_items ci 
			ON ci.product_id = p.id AND ci.user_id = $1
	`

	var products []models.Product
	err := r.db.Select(&products, query, idUser)
	if err != nil {
		return nil, response.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return products, nil
}

func (r *Repository) DecrementProduct(idUser string, idProduct string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Уменьшаем количество, если больше 1
	updateQuery := `
		UPDATE cart_items
		SET quantity = quantity - 1
		WHERE user_id = $1 AND product_id = $2 AND quantity >= 1;
	`
	if _, err := tx.Exec(updateQuery, idUser, idProduct); err != nil {
		return err
	}

	// Удаляем товар, если количество стало 1
	deleteQuery := `
		DELETE FROM cart_items
		WHERE user_id = $1 AND product_id = $2 AND quantity = 0;
	`
	if _, err := tx.Exec(deleteQuery, idUser, idProduct); err != nil {
		return err
	}

	return tx.Commit()
}
