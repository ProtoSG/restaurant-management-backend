package types

import "errors"

type InventoryPrice struct {
	Value float32 `json:"value"`
}

func NewInventoryPrice(value float32) (*InventoryPrice, error) {
	inventoryPrice := InventoryPrice{Value: value}
	if err := inventoryPrice.ensureIsValid(); err != nil {
		return nil, err
	}
	return &inventoryPrice, nil
}

func (this InventoryPrice) ensureIsValid() error {
	if this.Value < 0 {
		return errors.New("El Precio tiene que ser mayor o igual a 0")
	}
	return nil
}
