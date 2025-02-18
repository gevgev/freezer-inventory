package api

import (
	"github.com/gevgev/freezer-inventory/internal/api/handlers"
	"github.com/gevgev/freezer-inventory/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)
	itemHandler := handlers.NewItemHandler(db)
	inventoryHandler := handlers.NewInventoryHandler(db)
	categoryHandler := handlers.NewCategoryHandler(db)
	tagHandler := handlers.NewTagHandler(db)

	// Rate limiting middleware
	router.Use(middleware.RateLimit())

	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
	}

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		// Items
		api.GET("/items", itemHandler.List)
		api.POST("/items", itemHandler.Create)
		api.GET("/items/:id", itemHandler.Get)
		api.PUT("/items/:id", itemHandler.Update)
		api.DELETE("/items/:id", itemHandler.Delete)
		api.GET("/items/search", itemHandler.Search)

		// Inventory
		api.GET("/inventory/:item_id/status", inventoryHandler.GetStatus)
		api.GET("/inventory/:item_id/history", inventoryHandler.GetHistory)
		api.POST("/inventory", inventoryHandler.AddEntry)

		// Categories
		api.GET("/categories", categoryHandler.List)
		api.POST("/categories", middleware.AdminRequired(), categoryHandler.Create)

		// Tags
		api.GET("/tags", tagHandler.List)
		api.POST("/tags", middleware.AdminRequired(), tagHandler.Create)
	}

	return router
}
