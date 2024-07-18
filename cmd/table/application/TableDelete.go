package application

import (
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/repository"
	"restaurant-management-backend/cmd/table/domain/types"
)

type TableDelete struct {
	repository repository.TableRepository
}

func NewTableDelete(reposiroty repository.TableRepository) *TableDelete {
	return &TableDelete{repository: reposiroty}
}

func (this TableDelete) Execute(id int) error {
	tableId, err := types.NewTableId(id)
	if err != nil {
		return err
	}

	if table, _ := this.repository.GetById(tableId); table == nil {
		return domain.NewTableNotFound(*tableId)
	}

	return this.repository.Delete(tableId)
}
