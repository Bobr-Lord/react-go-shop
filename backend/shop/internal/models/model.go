package models

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ImageUrl    string `json:"imageUrl" binding:"required"`
}
type CreateProductResponse struct {
	Response string `json:"response"`
}
