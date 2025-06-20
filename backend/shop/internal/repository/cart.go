package repository

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"
)

func (r *Repository) AddCartItem(idItem string, idUser string) error {
	query := `
INSERT INTO cart_items (user_id, product_id, quantity)
VALUES ($1, $2, 1)
ON CONFLICT (user_id, product_id)
DO UPDATE SET quantity = cart_items.quantity + 1;
`
	_, err := r.db.Exec(query, idUser, idItem)
	return err

}

func (r *Repository) DeleteCartItem(idItem string, idUser string) error {
	query := `
		DELETE FROM cart_items
		WHERE user_id = $1 AND product_id = $2;
	`
	_, err := r.db.Exec(query, idUser, idItem)
	return err
}

func (r *Repository) GetCartItems(idUser string) ([]models.ProductWithQuantity, error) {
	query := `
		SELECT p.id, p.name, p.description, p.price, p.category, p.image_url, ci.quantity
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		WHERE ci.user_id = $1;
	`

	rows, err := r.db.Query(query, idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.ProductWithQuantity

	for rows.Next() {
		var item models.ProductWithQuantity
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Category,
			&item.ImageUrl,
			&item.Quantity,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
