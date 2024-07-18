package types

import "errors"

type InventoryQuantity struct {
	Value int `json:"value"`
}

func NewInventoryQuantity(value int) (*InventoryQuantity, error) {
	inventoryQuantity := InventoryQuantity{Value: value}
	if err := inventoryQuantity.ensureIsValid(); err != nil {
		return nil, err
	}
	return &inventoryQuantity, nil
}

func (this InventoryQuantity) ensureIsValid() error {
	if this.Value < 0 {
		return errors.New("La Cantidad debe ser mayor o igual a 0")
	}
	return nil
}
