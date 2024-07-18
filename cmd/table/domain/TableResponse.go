package domain

import (
	"restaurant-management-backend/cmd/table/domain/types"
	"restaurant-management-backend/cmd/table_category/domain"
)

type TableResponse struct {
	Id       *types.TableId        `json:"id"`
	Name     *types.TableName      `json:"name"`
	Category *domain.TableCategory `json:"category"`
	Status   *types.TableStatus    `json:"status"`
}

func NewTableResponse(id *types.TableId, name *types.TableName, category *domain.TableCategory, status *types.TableStatus) *TableResponse {
	return &TableResponse{
		Id:       id,
		Name:     name,
		Category: category,
		Status:   status,
	}
}
