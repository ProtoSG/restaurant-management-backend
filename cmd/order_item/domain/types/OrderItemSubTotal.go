package types

import "errors"

type OrderItemSubTotal struct {
	Value float32 `json:"value"`
}

func NewOrderSubTotal(value float32) (*OrderItemSubTotal, error) {
	orderSubTotal := OrderItemSubTotal{Value: value}
	if err := orderSubTotal.ensureIsValid(); err != nil {
		return nil, err
	}
	return &orderSubTotal, nil
}

func (this OrderItemSubTotal) ensureIsValid() error {
	if this.Value < 0 {
		return errors.New("El SubTotal no puede ser menor a 0")
	}
	return nil
}
