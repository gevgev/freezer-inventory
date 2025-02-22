package models

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	ItemID      uuid.UUID `json:"item_id" gorm:"type:uuid;primary_key"`
	ItemName    string    `json:"item_name" gorm:"->"` // Read-only field from join
	Quantity    int       `json:"quantity"`
	LastUpdated time.Time `json:"last_updated" gorm:"autoUpdateTime"`
	Weight      float64   `json:"weight"`
	WeightUnit  string    `json:"weight_unit"`
}
