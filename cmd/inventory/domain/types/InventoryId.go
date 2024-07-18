package types

type InventoryId struct {
	Value int `json:"value"`
}

func NewInventoryId(value int) (*InventoryId, error) {
	inventoryId := InventoryId{Value: value}
	return &inventoryId, nil
}

func (this InventoryId) ToValue() int {
	return this.Value
}
