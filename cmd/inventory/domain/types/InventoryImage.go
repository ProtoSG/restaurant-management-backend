package types

type InventoryImage struct {
	Value string `json:"value"`
}

func NewInventoryImage(value string) (*InventoryImage, error) {
	return &InventoryImage{Value: value}, nil
}
