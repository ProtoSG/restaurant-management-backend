package domain

type UserValidationError struct {
	Field   string
	Message string
}

func (this *UserPrimitive) Validate() *UserValidationError {
	if this.Username == "" {
		return &UserValidationError{Field: "username", Message: "Username is required"}
	}
	if this.Role == "" {
		return &UserValidationError{Field: "role", Message: "Role is required"}
	}
	if this.Password == "" {
		return &UserValidationError{Field: "password", Message: "Password is required"}
	}
	return nil
}
