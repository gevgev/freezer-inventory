package repository

import (
	"github.com/gevgev/freezer-inventory/internal/models"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	GetAllCurrentInventory() ([]models.Inventory, error)
}

type InventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{db: db}
}

func (r *InventoryRepositoryImpl) GetAllCurrentInventory() ([]models.Inventory, error) {
	var inventories []models.Inventory
	err := r.db.Table("items").
		Select(`
			items.id as item_id,
			items.name as item_name,
			COALESCE(SUM(inventory_log.change), 0) as quantity,
			MAX(inventory_log.timestamp) as last_updated,
			COALESCE(SUM(inventory_log.weight), 0) as weight,
			COALESCE(MAX(CASE WHEN inventory_log.change > 0 THEN inventory_log.weight_unit END), 'g') as weight_unit
		`).
		Joins("LEFT JOIN inventory_log ON items.id = inventory_log.item_id").
		Group("items.id, items.name").
		Having("COALESCE(SUM(inventory_log.change), 0) > 0").
		Find(&inventories).Error
	return inventories, err
}
