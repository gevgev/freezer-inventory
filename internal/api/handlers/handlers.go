package handlers

type Handlers struct {
	User      *UserHandler
	Inventory *InventoryHandler
}

func NewHandlers(userHandler *UserHandler, inventoryHandler *InventoryHandler) *Handlers {
	return &Handlers{
		User:      userHandler,
		Inventory: inventoryHandler,
	}
}
