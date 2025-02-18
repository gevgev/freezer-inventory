package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"type:user_role;not null"`
	CreatedAt    time.Time
}

type Item struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name           string    `gorm:"not null"`
	Description    string
	Barcode        string
	ImageURL       string
	Packaging      string
	WeightUnit     string `gorm:"type:weight_unit"`
	ExpirationDate time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Categories     []Category `gorm:"many2many:item_categories;"`
	Tags           []Tag      `gorm:"many2many:item_tags;"`
}

type InventoryLog struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ItemID     uuid.UUID `gorm:"type:uuid;not null"`
	Timestamp  time.Time `gorm:"not null"`
	Change     int       `gorm:"not null"`
	Weight     float64
	WeightUnit string `gorm:"type:weight_unit"`
	Notes      string
}

func (InventoryLog) TableName() string {
	return "inventory_log"
}

type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `gorm:"not null"`
	Description string
	Items       []Item `gorm:"many2many:item_categories;"`
}

type Tag struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name  string    `gorm:"not null"`
	Items []Item    `gorm:"many2many:item_tags;"`
}
