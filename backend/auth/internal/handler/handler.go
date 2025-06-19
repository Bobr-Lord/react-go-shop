package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/middleware"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "http://192.168.1.69:3000", "http://192.168.1.69:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(middleware.LoggerMiddleware())

	api := r.Group("/api")
	{
		api.POST("/reg", h.Register)
		api.POST("/login", h.Login)
	}

	return r
}
