package domain

import "restaurant-management-backend/cmd/table_category/domain"

type TableResponsePrimitive struct {
	Id       int                            `json:"id"`
	Name     string                         `json:"name"`
	Category *domain.TableCategoryPrimitive `json:"category"`
	Status   int                            `json:"status"`
}

func (this TableResponse) MapToPrimitive() *TableResponsePrimitive {
	return &TableResponsePrimitive{
		Id:       this.Id.Value,
		Name:     this.Name.Value,
		Category: this.Category.MapToPrimitive(),
		Status:   this.Status.Value,
	}
}
