package handler

import "github.com/gin-gonic/gin"

func (h *Handler) CreateProduct(c *gin.Context) {
	c.JSON(200, gin.H{})
}
