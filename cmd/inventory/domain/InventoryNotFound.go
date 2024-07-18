package domain

import (
	"fmt"
	"restaurant-management-backend/cmd/inventory/domain/types"
)

type InventoryNotFound struct {
	InventoryId *types.InventoryId
}

func NewInventoryNotFound(inventoryId *types.InventoryId) *InventoryNotFound {
	return &InventoryNotFound{InventoryId: inventoryId}
}

func (this InventoryNotFound) Error() string {
	return fmt.Sprintf("Inventario con el ID %d no encontrado", this.InventoryId)
}
