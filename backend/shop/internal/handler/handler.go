package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/config"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/middleware"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/service"
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
	r.Static("/static", "./frontend/build/static")
	r.StaticFile("/favicon.ico", "./frontend/build/favicon.ico")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.1.69:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.LoggerMiddleware())
	api := r.Group("/api")
	{
		api.POST("/product", middleware.AuthAdminMiddleware(h.cfg.PathPublicKey), h.CreateProduct)
		api.GET("/products", middleware.AuthUserMiddleware(h.cfg.PathPublicKey), h.GetAllProducts)
		api.DELETE("/product/:id", middleware.AuthAdminMiddleware(h.cfg.PathPublicKey), h.DeleteProduct)
		api.GET("/me", middleware.AuthUserMiddleware(h.cfg.PathPublicKey), h.GetMe)
	}

	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})
	r.GET("/admin", func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	return r
}
