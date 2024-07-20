package infrastructure

import (
	"database/sql"
	"restaurant-management-backend/cmd/order_item/application"
	"restaurant-management-backend/cmd/order_item/infrastructure"
)

type OrderItemService struct {
	Create     *application.OrderItemCreate
	GetAll     *application.OrderItemGetAll
	GetById    *application.OrderItemGetById
	Edit       *application.OrderItemEdit
	Delete     *application.OrderItemDelete
	GetByOrder *application.OrderItemGetByOrder
}

func NewOrderItemService(db *sql.DB) OrderItemService {
	orderItemContainer := infrastructure.NewSQLiteOrderItemRepository(db)

	return OrderItemService{
		Create:     application.NewOrderItemCreate(orderItemContainer),
		GetAll:     application.NewOrderItemGetAll(orderItemContainer),
		GetById:    application.NewOrderItemGetById(orderItemContainer),
		Edit:       application.NewOrderItemEdit(orderItemContainer),
		Delete:     application.NewOrderItemDeletee(orderItemContainer),
		GetByOrder: application.NewOrderItemGetByOrder(orderItemContainer),
	}
}
