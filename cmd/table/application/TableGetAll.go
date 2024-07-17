package application

import (
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/repository"
)

type TableGetAll struct {
	repository repository.TableRepository
}

func NewTableGetAll(repository repository.TableRepository) *TableGetAll {
	return &TableGetAll{repository: repository}
}

func (this TableGetAll) Execute() ([]*domain.Table, error) {
	return this.repository.GetAll()
}
