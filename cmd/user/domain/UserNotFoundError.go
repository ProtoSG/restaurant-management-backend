package domain

import (
	"fmt"
	"restaurant-management-backend/cmd/user/domain/types"
)

type UserNotFoundError struct {
	UserId types.UserId
}

func NewUserNotFoundError(userId types.UserId) *UserNotFoundError {
	return &UserNotFoundError{UserId: userId}
}

func (this *UserNotFoundError) Error() string {
	return fmt.Sprintf("Usuario con el ID %d no encontrado", this.UserId)
}
