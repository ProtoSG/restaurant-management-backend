package application

import (
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/order/domain/repository"
	"restaurant-management-backend/cmd/order/domain/types"
	"time"

	typesTable "restaurant-management-backend/cmd/table/domain/types"

	domainTable "restaurant-management-backend/cmd/table/domain"
	repositoryTable "restaurant-management-backend/cmd/table/domain/repository"
)

type OrderEdit struct {
	repository repository.OrderRepository
}

func NewOrderEdit(repository repository.OrderRepository) *OrderEdit {
	return &OrderEdit{
		repository: repository,
	}
}

func (this *OrderEdit) Execute(id int, tableId int, userId int, total float32, createdAt time.Time, updatedAt time.Time) error {
	orderId, err := types.NewOrderId(id)
	if err != nil {
		return err
	}

	if order, _ := this.repository.GetById(orderId); order == nil {
		return domain.NewOrderNotFound(orderId)
	}

	orderTableId, err := typesTable.NewTableId(tableId)
	if err != nil {
		return err
	}

	if table, _ := repositoryTable.TableRepository.GetById(this.repository, orderTableId); table == nil {
		return domainTable.NewTableNotFound(*orderTableId)
	}
}
