package types

import "errors"

type OrderCompleted struct {
	Value int `json:"value"`
}

func NewOrderCompleted(value int) (*OrderCompleted, error) {
	orderCompleted := OrderCompleted{Value: value}
	err := orderCompleted.ensureIsValid()
	if err != nil {
		return nil, err
	}
	return &orderCompleted, nil
}

func (this *OrderCompleted) ensureIsValid() error {
	if this.Value != 0 && this.Value != 1 {
		return errors.New("El valor del estado es incorrecto, permite solo 0 o 1")
	}
	return nil
}
