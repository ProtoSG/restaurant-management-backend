package infrastructure

import (
	"restaurant-management-backend/cmd/user/domain"
	"restaurant-management-backend/cmd/user/domain/types"
)

type InMemoryUserRepository struct {
	users []*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: []*domain.User{}}
}

func (this *InMemoryUserRepository) Create(user *domain.User) error {
	this.users = append(this.users, user)
	return nil
}

func (this *InMemoryUserRepository) GetAll() ([]*domain.User, error) {
	return this.users, nil
}

func (this *InMemoryUserRepository) GetById(id *types.UserId) (*domain.User, error) {
	for _, user := range this.users {
		if user.Id.ToValue() == id.ToValue() {
			return user, nil
		}
	}
	return nil, domain.NewUserNotFoundError(*id)
}

func (this *InMemoryUserRepository) Edit(user *domain.User) error {
	for i, u := range this.users {
		if u.Id.ToValue() == user.Id.ToValue() {
			this.users[i] = user
			return nil
		}
	}
	return domain.NewUserNotFoundError(*user.Id)
}

func (this *InMemoryUserRepository) Delete(id *types.UserId) error {
	for i, user := range this.users {
		if user.Id.ToValue() == id.ToValue() {
			this.users = append(this.users[:i], this.users[i+1:]...)
			return nil
		}
	}
	return domain.NewUserNotFoundError(*id)
}
