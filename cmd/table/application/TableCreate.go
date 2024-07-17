package application

import (
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/repository"
	"restaurant-management-backend/cmd/table/domain/types"
)

type TableCreate struct {
	repository repository.TableRepository
}

func NewTableCreate(repository repository.TableRepository) *TableCreate {
	return &TableCreate{repository: repository}
}

func (this TableCreate) Execute(id int, name string, category_id int, status int) error {
	tableId, err := types.NewTableId(id)
	if err != nil {
		return err
	}

	tableName, err := types.NewTableName(name)
	if err != nil {
		return err
	}

	tableCategoryId, err := types.NewTableCategoryId(category_id)
	if err != nil {
		return err
	}

	tableStatus, err := types.NewTableStatus(status)
	if err != nil {
		return err
	}

	table := domain.NewTable(tableId, tableName, tableCategoryId, tableStatus)
	return this.repository.Create(table)
}
