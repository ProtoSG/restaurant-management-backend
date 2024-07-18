package types

type InventoryName struct {
	Value string `json:"value"`
}

func NewInventoryName(value string) (*InventoryName, error) {
	inventoryName := InventoryName{Value: value}
	return &inventoryName, nil
}

func (this InventoryName) ToValue() string {
	return this.Value
}
