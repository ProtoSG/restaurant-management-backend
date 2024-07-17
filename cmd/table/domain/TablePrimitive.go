package domain

import "restaurant-management-backend/cmd/shared/domain"

type TablePrimitive struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
	Status     int    `json:"status"`
}

func (this Table) MapToPrimitive() *TablePrimitive {
	return &TablePrimitive{
		Id:         this.Id.Value,
		Name:       this.Name.Value,
		CategoryId: this.CategoyId.Value,
		Status:     this.Status.Value,
	}
}

func (this TablePrimitive) Validate() *domain.ValidationFieldError {
	if this.Name == "" {
		return &domain.ValidationFieldError{Field: "name", Message: "Name es requerido"}
	}
	if this.CategoryId == 0 {
		return &domain.ValidationFieldError{Field: "category_id", Message: "Category_id es requerido"}
	}
	if this.Status == 0 {
		return &domain.ValidationFieldError{Field: "status", Message: "Status es requerido"}
	}
	return nil
}
