package handlers

import (
	"net/http"
	"time"

	"log"

	"github.com/gevgev/freezer-inventory/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemHandler struct {
	db *gorm.DB
}

func NewItemHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{db: db}
}

func (h *ItemHandler) List(c *gin.Context) {
	var items []models.Item
	if err := h.db.Preload("Categories").
		Preload("Tags").
		Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}
	c.JSON(http.StatusOK, items)
}

type CreateItemRequest struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	Barcode        string `json:"barcode"`
	ImageURL       string `json:"image_url"`
	Packaging      string `json:"packaging"`
	WeightUnit     string `json:"weight_unit" binding:"required,oneof=kg g lb oz"`
	ExpirationDate string `json:"expiration_date" binding:"required"`
}

func (h *ItemHandler) Create(c *gin.Context) {
	var req CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expirationDate, err := time.Parse("2006-01-02", req.ExpirationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	item := models.Item{
		Name:           req.Name,
		Description:    req.Description,
		Barcode:        req.Barcode,
		ImageURL:       req.ImageURL,
		Packaging:      req.Packaging,
		WeightUnit:     req.WeightUnit,
		ExpirationDate: expirationDate,
	}

	if err := h.db.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var item models.Item
	if err := h.db.Preload("Categories").
		Preload("Tags").
		First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var item models.Item
	if err := h.db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.db.Delete(&models.Item{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func (h *ItemHandler) Search(c *gin.Context) {
	query := c.Query("q")
	var items []models.Item

	if err := h.db.Where("name ILIKE ?", "%"+query+"%").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

type ItemCategoryRequest struct {
	CategoryIDs []uuid.UUID `json:"category_ids" binding:"required"`
}

type ItemTagRequest struct {
	TagIDs []uuid.UUID `json:"tag_ids" binding:"required"`
}

func (h *ItemHandler) AddCategories(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var req ItemCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Load item
	var item models.Item
	if err := h.db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Load all categories to be added
	var categories []models.Category
	if err := h.db.Find(&categories, req.CategoryIDs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "One or more categories not found"})
		return
	}

	// Add the associations using GORM
	if err := h.db.Model(&item).Association("Categories").Append(&categories); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categories added successfully"})
}

func (h *ItemHandler) RemoveCategory(c *gin.Context) {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Invalid item ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	categoryID, err := uuid.Parse(c.Param("category_id"))
	if err != nil {
		log.Printf("Invalid category ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	log.Printf("Attempting to remove category %s from item %s", categoryID, itemID)

	// First verify both item and category exist
	var item models.Item
	if err := h.db.First(&item, itemID).Error; err != nil {
		log.Printf("Failed to find item %s: %v", itemID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	log.Printf("Found item: %s", item.Name)

	var category models.Category
	if err := h.db.First(&category, categoryID).Error; err != nil {
		log.Printf("Failed to find category %s: %v", categoryID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	log.Printf("Found category: %s", category.Name)

	// Check if the relationship exists
	var exists int64
	if err := h.db.Table("item_categories").
		Where("item_id = ? AND category_id = ?", itemID, categoryID).
		Count(&exists).Error; err != nil {
		log.Printf("Error checking relationship: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check relationship"})
		return
	}

	if exists == 0 {
		log.Printf("No relationship found between item %s and category %s", itemID, categoryID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Item is not associated with this category"})
		return
	}
	log.Printf("Found relationship between item and category")

	// Begin transaction
	tx := h.db.Begin()
	if tx.Error != nil {
		log.Printf("Failed to begin transaction: %v", tx.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Delete the relationship
	if err := tx.Table("item_categories").
		Where("item_id = ? AND category_id = ?", itemID, categoryID).
		Delete(&struct{}{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed to delete relationship: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove category"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit changes"})
		return
	}

	log.Printf("Successfully removed category %s from item %s", categoryID, itemID)
	c.JSON(http.StatusOK, gin.H{"message": "Category removed successfully"})
}

func (h *ItemHandler) AddTags(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var req ItemTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var item models.Item
	if err := h.db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := h.db.Model(&item).Association("Tags").Append(&models.Tag{ID: req.TagIDs[0]}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add tags"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tags added successfully"})
}

func (h *ItemHandler) RemoveTag(c *gin.Context) {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	tagID, err := uuid.Parse(c.Param("tag_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	if err := h.db.Exec("DELETE FROM item_tags WHERE item_id = ? AND tag_id = ?",
		itemID, tagID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag removed successfully"})
}
