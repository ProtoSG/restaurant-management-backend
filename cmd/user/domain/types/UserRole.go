package types

import "errors"

type UserRole struct {
	Value string `json:"value"`
}

func NewUserRole(value string) (*UserRole, error) {
	userRole := UserRole{Value: value}
	err := userRole.ensureIsValid()
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}

var validUserRoles = map[string]bool{
	"ADMIN":  true,
	"WAITER": true,
}

func (this UserRole) ensureIsValid() error {
	if _, ok := validUserRoles[this.Value]; !ok {
		return errors.New("Rol del User tiene que ser 'ADMIN' o 'WAITER'")
	}
	return nil
}

func (this UserRole) ToString() string {
	return this.Value
}
