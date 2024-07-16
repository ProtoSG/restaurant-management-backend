package types

import "errors"

type UserPassword struct {
	Value string `json:"value"`
}

func NewUserPassword(value string) (*UserPassword, error) {
	userPassword := UserPassword{Value: value}
	return &userPassword, nil
}

func (this UserPassword) ensureIsValid() error {
	if len(this.Value) < 5 {
		return errors.New("Password del User tiene que tener mas de 5 caracteres")
	}
	return nil
}

func (this UserPassword) ToString() string {
	return this.Value
}
