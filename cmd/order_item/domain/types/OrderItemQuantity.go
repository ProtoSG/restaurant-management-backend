package types

import "errors"

type OrderItemQuantity struct {
	Value int `json:"value"`
}

func NewOrderQuantity(value int) (*OrderItemQuantity, error) {
	orderQuantity := OrderItemQuantity{Value: value}
	if err := orderQuantity.ensureIsValid(); err != nil {
		return nil, err
	}
	return &orderQuantity, nil
}

func (this OrderItemQuantity) ensureIsValid() error {
	if this.Value <= 0 {
		return errors.New("La cantidad tiene que ser mayor a 0")
	}
	return nil
}
