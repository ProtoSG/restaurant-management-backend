package domain

import "restaurant-management-backend/cmd/shared/domain"

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

func (this *UserPrimitive) Validate() *domain.ValidationFieldError {
	if this.Username == "" {
		return &domain.ValidationFieldError{Field: "username", Message: "Username is required"}
	}
	if this.Role == "" {
		return &domain.ValidationFieldError{Field: "role", Message: "Role is required"}
	}
	if this.Password == "" {
		return &domain.ValidationFieldError{Field: "password", Message: "Password is required"}
	}
	return nil
}
