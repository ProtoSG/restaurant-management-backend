package application

import (
	"restaurant-management-backend/cmd/user/domain/repository"
	"restaurant-management-backend/cmd/user/domain/types"
)

type UserDelete struct {
	repository repository.UserRepository
}

func NewUserDelete(repository repository.UserRepository) *UserDelete {
	return &UserDelete{repository: repository}
}

func (this *UserDelete) Execute(id int) error {
	userId, err := types.NewUserId(id)
	if err != nil {
		return err
	}

	return this.repository.Delete(userId)
}
