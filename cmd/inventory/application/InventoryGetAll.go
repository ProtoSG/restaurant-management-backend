package application

import (
	"restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/inventory/domain/repository"
)

type InventoryGetAll struct {
	repository repository.InventoryRepository
}

func NewInventoryGetAll(repository repository.InventoryRepository) *InventoryGetAll {
	return &InventoryGetAll{
		repository: repository,
	}
}

func (this *InventoryGetAll) Execute() ([]*domain.InventoryResponse, error) {
	return this.repository.GetAll()
}
