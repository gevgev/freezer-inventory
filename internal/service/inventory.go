package service

import (
	"github.com/gevgev/freezer-inventory/internal/models"
	"github.com/gevgev/freezer-inventory/internal/repository"
)

type InventoryService interface {
	GetAllCurrentInventory() ([]models.Inventory, error)
}

type InventoryServiceImpl struct {
	repo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) *InventoryServiceImpl {
	return &InventoryServiceImpl{
		repo: repo,
	}
}

func (s *InventoryServiceImpl) GetAllCurrentInventory() ([]models.Inventory, error) {
	return s.repo.GetAllCurrentInventory()
}
