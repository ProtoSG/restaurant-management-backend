package application

import (
	"restaurant-management-backend/cmd/user/domain"
	"restaurant-management-backend/cmd/user/domain/repository"
)

type UserGetAll struct {
	repository repository.UserRepository
}

func NewUserGetAll(repository repository.UserRepository) *UserGetAll {
	return &UserGetAll{repository: repository}
}

func (this *UserGetAll) Execute() ([]*domain.User, error) {
	return this.repository.GetAll()
}
