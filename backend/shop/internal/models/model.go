package models

type Product struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	Category    string `json:"category" db:"category"`
	ImageUrl    string `json:"image" db:"image_url"`
	Quantity    int    `json:"quantity" db:"quantity"`
}

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ImageUrl    string `json:"image" binding:"required"`
}
type CreateProductResponse struct {
	ID string `json:"id"`
}

type GetAllProductsRequest struct {
}
type GetAllProductsResponse struct {
	Products []Product `json:"products"`
}

type DeleteProductRequest struct {
}
type DeleteProductResponse struct {
	Response string `json:"response"`
}

type AddCartItemRequest struct {
	IDItem string `json:"id" binding:"required"`
}

type DeleteCartItemRequest struct {
	IDItem string `json:"id" binding:"required"`
}

type GetCartItemsResponse struct {
	Products []Product `json:"products"`
}

type ProductWithQuantity struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ImageUrl    string `json:"image" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}
