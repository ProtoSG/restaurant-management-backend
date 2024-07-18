package domain

import "restaurant-management-backend/cmd/shared/domain"

type ItemCategoryPrimitive struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (this ItemCategory) MapToPrimitive() *ItemCategoryPrimitive {
	return &ItemCategoryPrimitive{
		Id:   this.Id.Value,
		Name: this.Name.Value,
	}
}

func (this *ItemCategoryPrimitive) Validate() *domain.ValidationFieldError {
	if this.Name == "" {
		return &domain.ValidationFieldError{Field: "name", Message: "Name is required"}
	}
	return nil
}
