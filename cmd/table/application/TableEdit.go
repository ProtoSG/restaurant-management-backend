package application

import (
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/repository"
	"restaurant-management-backend/cmd/table/domain/types"
)

type TableEdit struct {
	repository repository.TableRepository
}

func NewTableEdit(repository repository.TableRepository) *TableEdit {
	return &TableEdit{repository: repository}
}

func (this TableEdit) Execute(id int, name string, category_id int, status int) error {
	tableId, err := types.NewTableId(id)
	if err != nil {
		return err
	}

	if table, _ := this.repository.GetById(tableId); table == nil {
		return domain.NewTableNotFound(*tableId)
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

	return this.repository.Edit(table)
}
