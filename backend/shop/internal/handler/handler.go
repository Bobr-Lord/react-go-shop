package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (Handler) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", func(c *gin.Context) {})

	return r
}
