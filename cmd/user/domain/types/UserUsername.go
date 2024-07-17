package types

import "errors"

type UserUsername struct {
	Value string `json:"value"`
}

func NewUserUsername(value string) (*UserUsername, error) {
	userUsername := UserUsername{Value: value}
	// if err := userUsername.ensureIsValid(); err != nil {
	// 	return nil, err
	// }
	return &userUsername, nil
}

func (this UserUsername) ensureIsValid() error {
	if len(this.Value) < 5 {
		return errors.New("Username tiene que tener mas de 5 carácteres")
	}
	return nil
}

func (this UserUsername) ToString() string {
	return this.Value
}
