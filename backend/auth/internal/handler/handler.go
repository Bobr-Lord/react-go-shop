package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/config"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/middleware"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	svc *service.Service
	cfg *config.Config
}

func NewHandler(svc *service.Service, cfg *config.Config) *Handler {
	return &Handler{
		svc: svc,
		cfg: cfg,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.LoggerMiddleware())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.POST("/reg", h.Register)
		api.POST("/login", h.Login)
		api.GET("/me", middleware.AuthUserMiddleware(h.cfg.PathPublicKey), h.GetMe)
		api.GET("/verify", h.VerifyEmail)
	}

	return r
}
