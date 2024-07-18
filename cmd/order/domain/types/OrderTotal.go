package types

import "errors"

type OrderTotal struct {
	Value float32 `json:"value"`
}

func NewOrderTotal(value float32) (*OrderTotal, error) {
	orderTotal := OrderTotal{Value: value}
	return &orderTotal, nil
}

func (this OrderTotal) ensureIsValid() error {
	if this.Value < 0 {
		return errors.New("El total de la orden tiene que ser mayor o igual a 0")
	}
	return nil
}
