package infrastructure

import (
	"database/sql"
	"fmt"
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/order/domain/types"

	domainUser "restaurant-management-backend/cmd/user/domain"
	typesUser "restaurant-management-backend/cmd/user/domain/types"

	domainTable "restaurant-management-backend/cmd/table/domain"
	typesTable "restaurant-management-backend/cmd/table/domain/types"
)

type SQLiteOrderRepository struct {
	db *sql.DB
}

func NewSQLiteOrderRepository(db *sql.DB) *SQLiteOrderRepository {
	return &SQLiteOrderRepository{db: db}
}

func (this SQLiteOrderRepository) Create(order *domain.Order) error {
	stmt, err := this.db.Prepare("INSERT INTO orders (table_id, user_id, total, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.TableId.Value, order.UserId.Value, order.Total.Value, order.CreatedAt.Value, order.UpdatedAt.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteOrderRepository) GetAll() ([]*domain.OrderResponse, error) {
	stmt, err := this.db.Prepare(`
    SELECT * FROM orders;
  `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.OrderResponse

	for rows.Next() {
		var order domain.OrderPrimitive
		var table *domainTable.TableResponse
		var user *domainUser.User

		if err := rows.Scan(&order.Id, &order.TableId, &order.UserId, &order.Total, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}

		orderId, _ := types.NewOrderId(order.Id)
		orderTableId, _ := typesTable.NewTableId(order.TableId)
		orderUserId, _ := typesUser.NewUserId(order.UserId)
		orderTotal, _ := types.NewOrderTotal(order.Total)
		orderCreatedAt, _ := types.NewOrderCreatedAt(order.CreatedAt)
		orderUpdatedAt, _ := types.NewOrderUpdatedAt(order.UpdatedAt)

		newOrder := domain.NewOrderResponse(orderId, orderTableId, table, orderUserId, user, orderTotal, orderCreatedAt, orderUpdatedAt)

		orders = append(orders, newOrder)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration; ", err)
	}

	return orders, nil
}

func (this SQLiteOrderRepository) GetById(orderId *types.OrderId) (*domain.OrderResponse, error) {
	stmt, err := this.db.Prepare("SELECT * FROM orders WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(orderId.Value)

	var order domain.OrderPrimitive
	var user *domainUser.User
	var table *domainTable.TableResponse

	err = row.Scan(&order.Id, &order.TableId, &order.UserId, &order.Total, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	orderTableId, _ := typesTable.NewTableId(order.TableId)
	orderUserId, _ := typesUser.NewUserId(order.UserId)
	orderTotal, _ := types.NewOrderTotal(order.Total)
	orderCreatedAt, _ := types.NewOrderCreatedAt(order.CreatedAt)
	orderUpdatedAt, _ := types.NewOrderUpdatedAt(order.UpdatedAt)

	newOrder := domain.NewOrderResponse(orderId, orderTableId, table, orderUserId, user, orderTotal, orderCreatedAt, orderUpdatedAt)

	return newOrder, nil
}

func (this SQLiteOrderRepository) Edit(order *domain.Order) error {
	stmt, err := this.db.Prepare("UPDATE orders SET table_id = ?, user_id = ?, total = ?, created_at = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.TableId, order.UserId, order.Total, order.CreatedAt, order.UpdatedAt, order.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteOrderRepository) Delete(orderId *types.OrderId) error {
	stmt, err := this.db.Prepare("DELETE FROM tables WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderId.Value)
	if err != nil {
		return err
	}

	return nil
}
