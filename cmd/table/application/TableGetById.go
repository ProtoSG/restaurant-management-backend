package application

import (
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/repository"
	"restaurant-management-backend/cmd/table/domain/types"
)

type TableGetById struct {
	repository repository.TableRepository
}

func NewTableGetById(repository repository.TableRepository) *TableGetById {
	return &TableGetById{repository: repository}
}

func (this TableGetById) Execute(id int) (*domain.Table, error) {
	tableId, err := types.NewTableId(id)
	if err != nil {
		return nil, err
	}

	table, err := this.repository.GetById(tableId)
	if err != nil {
		return nil, domain.NewTableNotFound(*tableId)
	}

	return table, nil
}
