package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/middleware"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/service"
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

func (h *Handler) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.LoggerMiddleware())

	return r
}
