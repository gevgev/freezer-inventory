package repository

import (
	"github.com/gevgev/freezer-inventory/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.InventoryLog{},
		&models.Category{},
		&models.Tag{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
