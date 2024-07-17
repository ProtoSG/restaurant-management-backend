package domain

import "restaurant-management-backend/cmd/shared/domain"

type TableCategoryPrimitive struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (this TableCategory) MapToPrimitive() *TableCategoryPrimitive {
	return &TableCategoryPrimitive{
		Id:   this.Id.Value,
		Name: this.Name.Value,
	}
}

func (this TableCategoryPrimitive) Validate() *domain.ValidationFieldError {
	if this.Name == "" {
		return &domain.ValidationFieldError{Field: "name", Message: "Name es requerido"}
	}
	return nil
}
