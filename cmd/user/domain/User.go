package domain

import "restaurant-management-backend/cmd/user/domain/types"

type User struct {
	Id       types.UserId       `json:"id"`
	Username types.UserUsername `json:"username"`
	Password types.UserPassword `json:"password"`
	Role     types.UserRole     `json:"role"`
}

func NewUser(id types.UserId, username types.UserUsername, password types.UserPassword, role types.UserRole) *User {
	return &User{
		Id:       id,
		Username: username,
		Password: password,
		Role:     role,
	}
}

type UserPrimitive struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (this User) MapToPrimitive() *UserPrimitive {
	return &UserPrimitive{
		Id:       this.Id.Value,
		Username: this.Username.Value,
		Password: this.Password.Value,
		Role:     this.Role.Value,
	}
}
