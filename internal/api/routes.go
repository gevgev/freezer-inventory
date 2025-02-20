package api

import (
	"time"

	"github.com/gevgev/freezer-inventory/internal/api/handlers"
	"github.com/gevgev/freezer-inventory/internal/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// CORS middleware configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Your React app URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)
	userHandler := handlers.NewUserHandler(db)
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
		// User Management (Admin only)
		users := api.Group("/users")
		users.Use(middleware.AdminRequired())
		{
			users.GET("", userHandler.List)
			users.POST("", userHandler.Create)
			users.GET("/:id", userHandler.Get)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}

		// Items
		items := api.Group("/items")
		{
			items.GET("", itemHandler.List)
			items.POST("", itemHandler.Create)
			items.GET("/:id", itemHandler.Get)
			items.PUT("/:id", itemHandler.Update)
			items.DELETE("/:id", itemHandler.Delete)
			items.GET("/search", itemHandler.Search)
			items.POST("/:id/categories", itemHandler.AddCategories)
			items.DELETE("/:id/categories/:category_id", itemHandler.RemoveCategory)
			items.POST("/:id/tags", itemHandler.AddTags)
			items.DELETE("/:id/tags/:tag_id", itemHandler.RemoveTag)
		}

		// Inventory
		api.GET("/inventory/:item_id/status", inventoryHandler.GetStatus)
		api.GET("/inventory/:item_id/history", inventoryHandler.GetHistory)
		api.POST("/inventory", inventoryHandler.AddEntry)

		// Categories
		api.GET("/categories", categoryHandler.List)
		api.POST("/categories", middleware.AdminRequired(), categoryHandler.Create)
		api.PUT("/categories/:id", middleware.AdminRequired(), categoryHandler.Update)

		// Tags
		api.GET("/tags", tagHandler.List)
		api.POST("/tags", middleware.AdminRequired(), tagHandler.Create)
		api.PUT("/tags/:id", middleware.AdminRequired(), tagHandler.Update)
	}

	return router
}
