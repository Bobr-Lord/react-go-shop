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
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(middleware.LoggerMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		shop := api.Group("/shop")
		{
			shop.POST("/product", middleware.AuthAdminMiddleware(h.cfg.PathPublicKey), h.CreateProduct)
			shop.GET("/products/cart", middleware.AuthUserMiddleware(h.cfg.PathPublicKey), h.GetAllProductsPrivate)
			shop.GET("/products", h.GetAllProducts)
			shop.DELETE("/product/:id", middleware.AuthAdminMiddleware(h.cfg.PathPublicKey), h.DeleteProduct)

			cart := shop.Group("/cart")
			cart.Use(middleware.AuthUserMiddleware(h.cfg.PathPublicKey))
			{
				cart.GET("/item", h.GetCartItems)
				cart.POST("/item", h.AddCartItem)
				cart.DELETE("/item/:id_item", h.DeleteCartItem)
				cart.PUT("/item/:id", h.DecrementProduct)
			}
		}
	}
	//r.Static("/static", "./frontend/build/static")
	//r.StaticFile("/favicon.ico", "./frontend/build/favicon.ico")
	//r.StaticFile("/logo192.png", "./frontend/build/logo192.png") // по желанию
	//
	//r.NoRoute(func(c *gin.Context) {
	//	c.File("./frontend/build/index.html")
	//})
	return r
}
