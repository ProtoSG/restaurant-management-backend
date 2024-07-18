package types

type InventoryItemCategoryId struct {
	Value int `json:"value"`
}

func NewInventoryItemCategoryId(value int) (*InventoryItemCategoryId, error) {
	inventoryItemCategoryId := InventoryItemCategoryId{Value: value}
	return &inventoryItemCategoryId, nil
}

func (this InventoryItemCategoryId) ToValue() int {
	return this.Value
}
