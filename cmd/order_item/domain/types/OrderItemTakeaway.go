package types

import "errors"

type OrderItemTakeaway struct {
	Value int `json:"value"`
}

func NewOrderItemTakeaway(value int) (*OrderItemTakeaway, error) {
	orderItemTakeaway := OrderItemTakeaway{Value: value}
	if err := orderItemTakeaway.ensureIsValid(); err != nil {
		return nil, err
	}
	return &orderItemTakeaway, nil
}

func (this *OrderItemTakeaway) ensureIsValid() error {
	if this.Value != 0 && this.Value != 1 {
		return errors.New("El valor del estado es incorrecto, permite solo 0 o 1")
	}
	return nil
}
