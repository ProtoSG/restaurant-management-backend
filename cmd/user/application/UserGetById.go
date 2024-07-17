package application

import (
	"log"
	"restaurant-management-backend/cmd/user/domain"
	"restaurant-management-backend/cmd/user/domain/repository"
	"restaurant-management-backend/cmd/user/domain/types"
)

type UserGetById struct {
	repository repository.UserRepository
}

func NewUserGetById(repository repository.UserRepository) *UserGetById {
	return &UserGetById{repository: repository}
}

func (this *UserGetById) Execute(id int) (*domain.User, error) {
	userId, err := types.NewUserId(id)
	if err != nil {
		return nil, err
	}

	user, erro := this.repository.GetById(userId)
	if erro != nil {
		log.Println("Error", erro)
		return nil, domain.NewUserNotFoundError(*userId)
	}

	return user, nil
}
