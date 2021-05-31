package handler

import (
	"time"

	"github.com/daria40tim/plant-api/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func CORSM() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 120 * time.Hour,
	}))

	router.Use(CORSM())

	api := router.Group("/api")
	{
		ocr := api.Group("/ocr")
		{
			ocr.GET("/dirs", h.getDirs)
			ocr.POST("/img", h.getImg)
		}
		cards := api.Group("/cards")
		{
			cards.POST("/", h.createCard)
			cards.GET("/", h.getAllCards)
			cards.GET("/:id", h.getCardById)
			cards.PUT("/:id", h.updateCardById)
		}
		noms := api.Group("/noms")
		{
			noms.POST("/", h.createNom)
			noms.GET("/", h.getAllNoms)
			noms.GET("/:id", h.getNomById)
			noms.PUT("/:id", h.updateNomById)
		}
		orders := api.Group("/orders")
		{
			orders.POST("/", h.createOrder)
			orders.GET("/", h.getAllOrders)
			orders.GET("/:id", h.getOrderById)
			orders.PUT("/:id", h.updateOrderById)
		}
	}

	/*router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		/*		AllowMethods:     []string{"PUT", "POST", "GET"},
				AllowHeaders:     []string{"Origin"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
	}))*/

	return router
}
