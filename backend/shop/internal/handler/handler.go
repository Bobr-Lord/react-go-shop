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
		api.GET("/products/cart", middleware.AuthUserMiddleware(h.cfg.PathPublicKey), h.GetAllProductsPrivate)
		api.GET("/products", h.GetAllProducts)
		api.DELETE("/product/:id", middleware.AuthAdminMiddleware(h.cfg.PathPublicKey), h.DeleteProduct)

		cart := api.Group("/cart")
		cart.Use(middleware.AuthUserMiddleware(h.cfg.PathPublicKey))
		{
			cart.GET("/item", h.GetCartItems)
			cart.POST("/item", h.AddCartItem)
			cart.DELETE("/item/:id_item", h.DeleteCartItem)
			cart.PUT("/item/:id", h.DecrementProduct)
		}
	}
	r.Static("/static", "./frontend/build/static")
	r.StaticFile("/favicon.ico", "./frontend/build/favicon.ico")
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
