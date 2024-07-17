package types

import "errors"

type TableStatus struct {
	Value int `json:"value"`
}

func NewTableStatus(value int) (*TableStatus, error) {
	tableStatus := TableStatus{Value: value}
	if err := tableStatus.ensureIsValid(); err != nil {
		return nil, err
	}
	return &tableStatus, nil
}

func (this TableStatus) ensureIsValid() error {
	if this.Value != 0 && this.Value != 1 {
		return errors.New("El valor del stado es incorrecto, permite solo 0 o 1")
	}
	return nil
}

func (this TableStatus) ToValue() int {
	return this.Value
}
