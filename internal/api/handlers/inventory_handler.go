package handlers

import (
	"net/http"
	"time"

	"github.com/gevgev/freezer-inventory/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryHandler struct {
	db *gorm.DB
}

func NewInventoryHandler(db *gorm.DB) *InventoryHandler {
	return &InventoryHandler{db: db}
}

func (h *InventoryHandler) GetStatus(c *gin.Context) {
	itemID, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var total int64
	if err := h.db.Model(&models.InventoryLog{}).
		Where("item_id = ?", itemID).
		Select("COALESCE(SUM(change), 0)").
		Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inventory status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total})
}

func (h *InventoryHandler) GetHistory(c *gin.Context) {
	itemID, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var logs []models.InventoryLog
	if err := h.db.Where("item_id = ?", itemID).
		Order("timestamp desc").
		Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inventory history"})
		return
	}

	c.JSON(http.StatusOK, logs)
}

type AddInventoryRequest struct {
	ItemID     uuid.UUID `json:"item_id" binding:"required"`
	Change     int       `json:"change" binding:"required"`
	Weight     float64   `json:"weight" binding:"required"`
	WeightUnit string    `json:"weight_unit" binding:"required,oneof=kg g lb oz"`
	Notes      string    `json:"notes"`
}

func (h *InventoryHandler) AddEntry(c *gin.Context) {
	var req AddInventoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log := models.InventoryLog{
		ItemID:     req.ItemID,
		Timestamp:  time.Now(),
		Change:     req.Change,
		Weight:     req.Weight,
		WeightUnit: req.WeightUnit,
		Notes:      req.Notes,
	}

	if err := h.db.Create(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add inventory entry"})
		return
	}

	c.JSON(http.StatusCreated, log)
}
