package service

type Services struct {
	Inventory InventoryService
}

func NewServices(inventory InventoryService) *Services {
	return &Services{
		Inventory: inventory,
	}
}
