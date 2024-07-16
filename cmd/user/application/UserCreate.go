package application

import "restaurant-management-backend/cmd/user/domain/repository"

type UserCreate struct {
	repository repository.UserRepository
}
