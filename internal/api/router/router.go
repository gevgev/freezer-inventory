package router

import (
	"github.com/gevgev/freezer-inventory/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handlers *handlers.Handlers) {
	// ... existing routes ...

	// Add this new route
	router.GET("/api/v1/inventory", handlers.Inventory.GetCurrentInventory)

	// ... existing code ...
}
