package repository

import (
	"restaurant-management-backend/cmd/user/domain"
	"restaurant-management-backend/cmd/user/domain/types"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetAll() ([]*domain.User, error)
	GetById(id *types.UserId) (*domain.User, error)
	Edit(user *domain.User) error
	Delete(id *types.UserId) error
}
