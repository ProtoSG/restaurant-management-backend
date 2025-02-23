package application

import (
	"restaurant-management-backend/cmd/user/domain"
	"restaurant-management-backend/cmd/user/domain/repository"
	"restaurant-management-backend/cmd/user/domain/types"
)

type UserCreate struct {
	repository repository.UserRepository
}

func NewUserCreate(repository repository.UserRepository) *UserCreate {
	return &UserCreate{repository: repository}
}

func (this *UserCreate) Execute(id int, username string, password string, role string) error {
	userId, err := types.NewUserId(id)
	if err != nil {
		return err
	}

	userUsername, err := types.NewUserUsername(username)
	if err != nil {
		return err
	}

	userPassword, err := types.NewUserPassword(password)
	if err != nil {
		return err
	}

	userRole, err := types.NewUserRole(role)
	if err != nil {
		return err
	}

	user := domain.NewUser(userId, userUsername, userPassword, userRole)

	return this.repository.Create(user)
}
