package infrastructure

import (
	"database/sql"
	"restaurant-management-backend/cmd/order/application"
	"restaurant-management-backend/cmd/order/infrastructure"
)

type OrderService struct {
	Create  *application.OrderCreate
	GetAll  *application.OrderGetAll
	GetById *application.OrderGetById
	Edit    *application.OrderEdit
	Delete  *application.OrderDelete
}

func NewOrderService(db *sql.DB) OrderService {
	orderContainer := infrastructure.NewSQLiteOrderRepository(db)
	return OrderService{
		Create:  application.NewOrderCreate(orderContainer),
		GetAll:  application.NewOrderGetAll(orderContainer),
		GetById: application.NewOrderGetById(orderContainer),
		Edit:    application.NewOrderEdit(orderContainer),
		Delete:  application.NewOrderDelete(orderContainer),
	}
}
