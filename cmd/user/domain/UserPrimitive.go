package domain

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
