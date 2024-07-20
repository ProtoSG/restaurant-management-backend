package infrastructure

import (
	"database/sql"
	"fmt"
	domainInventory "restaurant-management-backend/cmd/inventory/domain"
	typesInventory "restaurant-management-backend/cmd/inventory/domain/types"
	typesOrder "restaurant-management-backend/cmd/order/domain/types"
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type SQLiteOrderItemRepository struct {
	db *sql.DB
}

func NewSQLiteOrderItemRepository(db *sql.DB) *SQLiteOrderItemRepository {
	return &SQLiteOrderItemRepository{db: db}
}

func (this SQLiteOrderItemRepository) Create(orderItem *domain.OrderItem) error {
	stmt, err := this.db.Prepare("INSERT INTO order_item (order_id, item_id, quantity, sub_total) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderItem.OrderId.Value, orderItem.ItemId.Value, orderItem.Quantity.Value, orderItem.SubTotal.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteOrderItemRepository) GetAll() ([]*domain.OrderItemResponse, error) {
	stmt, err := this.db.Prepare("SELECT * FROM order_item")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders_items []*domain.OrderItemResponse

	for rows.Next() {
		var order_item domain.OrderItemPrimitive
		var item *domainInventory.InventoryResponse

		if err := rows.Scan(&order_item.Id, &order_item.OrderId, &order_item.ItemId, &order_item.Quantity, &order_item.SubTotal); err != nil {
			return nil, err
		}

		orderItemId, _ := types.NewOrderItemId(order_item.Id)
		orderItemOrderId, _ := typesOrder.NewOrderId(order_item.OrderId)
		orderItemInventoryId, _ := typesInventory.NewInventoryId(order_item.ItemId)
		orderItemQuantity, _ := types.NewOrderQuantity(order_item.Quantity)
		orderItemSubTotal, _ := types.NewOrderSubTotal(order_item.SubTotal)

		newOrderItem := domain.NewOrderItemResponse(orderItemId, orderItemOrderId, orderItemInventoryId, item, orderItemQuantity, orderItemSubTotal)

		orders_items = append(orders_items, newOrderItem)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error durante la iteración de filas", err)
	}

	return orders_items, nil
}

func (this SQLiteOrderItemRepository) GetById(orderItemId *types.OrderItemId) (*domain.OrderItemResponse, error) {
	stmt, err := this.db.Prepare("SELECT * FROM order_item WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(orderItemId.Value)

	var order_item domain.OrderItemPrimitive
	var item *domainInventory.InventoryResponse

	if err = row.Scan(&order_item.Id, &order_item.OrderId, &order_item.ItemId, &order_item.Quantity, &order_item.SubTotal); err != nil {
		return nil, err
	}

	orderItemOrderId, _ := typesOrder.NewOrderId(order_item.OrderId)
	orderItemInventoryId, _ := typesInventory.NewInventoryId(order_item.ItemId)
	orderItemQuantity, _ := types.NewOrderQuantity(order_item.Quantity)
	orderItemSubTotal, _ := types.NewOrderSubTotal(order_item.SubTotal)

	newOrderItem := domain.NewOrderItemResponse(orderItemId, orderItemOrderId, orderItemInventoryId, item, orderItemQuantity, orderItemSubTotal)

	return newOrderItem, nil
}

func (this SQLiteOrderItemRepository) Edit(orderItem *domain.OrderItem) error {
	stmt, err := this.db.Prepare("UPDATE order_item SET order_id = ?, item_id = ?, quantity = ?, sub_total = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderItem.OrderId.Value, orderItem.ItemId.Value, orderItem.Quantity.Value, orderItem.SubTotal.Value, orderItem.Id.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteOrderItemRepository) Delete(orderItemId *types.OrderItemId) error {
	stmt, err := this.db.Prepare("DELETE FROM order_item WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(orderItemId.Value); err != nil {
		return err
	}

	return nil
}

func (this SQLiteOrderItemRepository) GetByOrder(orderId *typesOrder.OrderId) ([]*domain.OrderItemResponse, error) {
	stmt, err := this.db.Prepare("SELECT * FROM order_item WHERE order_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(orderId.Value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders_items []*domain.OrderItemResponse

	for rows.Next() {
		var order_item domain.OrderItemPrimitive
		var item *domainInventory.InventoryResponse

		if err := rows.Scan(&order_item.Id, &order_item.OrderId, &order_item.ItemId, &order_item.Quantity, &order_item.SubTotal); err != nil {
			return nil, err
		}

		orderItemId, _ := types.NewOrderItemId(order_item.Id)
		orderItemOrderId, _ := typesOrder.NewOrderId(order_item.OrderId)
		orderItemInventoryId, _ := typesInventory.NewInventoryId(order_item.ItemId)
		orderItemQuantity, _ := types.NewOrderQuantity(order_item.Quantity)
		orderItemSubTotal, _ := types.NewOrderSubTotal(order_item.SubTotal)

		newOrderItem := domain.NewOrderItemResponse(orderItemId, orderItemOrderId, orderItemInventoryId, item, orderItemQuantity, orderItemSubTotal)

		orders_items = append(orders_items, newOrderItem)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error durante la iteración de filas", err)
	}

	return orders_items, nil
}
