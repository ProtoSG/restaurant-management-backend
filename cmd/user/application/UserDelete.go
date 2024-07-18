package application

import (
	"restaurant-management-backend/cmd/user/domain"
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

	if user, _ := this.repository.GetById(userId); user == nil {
		return domain.NewUserNotFoundError(*userId)
	}

	return this.repository.Delete(userId)
}
