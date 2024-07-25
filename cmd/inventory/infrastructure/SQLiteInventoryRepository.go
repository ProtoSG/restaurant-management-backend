package infrastructure

import (
	"database/sql"
	"fmt"
	"restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/inventory/domain/types"
	domainItemCategory "restaurant-management-backend/cmd/item_category/domain"
	typesItemCategory "restaurant-management-backend/cmd/item_category/domain/types"
)

type SQLiteInventoryRepository struct {
	db *sql.DB
}

func NewSQLiteInventoryRepository(db *sql.DB) *SQLiteInventoryRepository {
	return &SQLiteInventoryRepository{db}
}

func (this SQLiteInventoryRepository) Create(inventory *domain.Inventory) error {
	stmt, err := this.db.Prepare("INSERT INTO inventory (name, item_category_id, quantity, price, image) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(inventory.Name.Value, inventory.ItemCategoryId.Value, inventory.Quantity.Value, inventory.Price.Value, inventory.Image.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteInventoryRepository) GetAll() ([]*domain.InventoryResponse, error) {
	stmt, err := this.db.Prepare(`
    SELECT i.id, i.name, i.item_category_id AS category_name, i.quantity, i.price, i.image, ic.name AS category_name
    FROM inventory i 
    JOIN item_category ic ON i.item_category_id = ic.id;
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

	var inventory []*domain.InventoryResponse
	for rows.Next() {
		var inv domain.InventoryPrimitive
		var itemCategory domainItemCategory.ItemCategoryPrimitive
		if err := rows.Scan(&inv.Id, &inv.Name, &itemCategory.Id, &inv.Quantity, &inv.Price, &inv.Image, &itemCategory.Name); err != nil {
			return nil, err
		}

		invId, _ := types.NewInventoryId(inv.Id)
		invName, _ := types.NewInventoryName(inv.Name)
		invQuantity, _ := types.NewInventoryQuantity(inv.Quantity)
		invPrice, _ := types.NewInventoryPrice(inv.Price)
		invImage, _ := types.NewInventoryImage(inv.Image)

		itemCategoryId, _ := typesItemCategory.NewItemCategoryId(itemCategory.Id)
		itemCategoryName, _ := typesItemCategory.NewItemCategoryName(itemCategory.Name)

		newItemCategory := domainItemCategory.NewItemCategory(itemCategoryId, itemCategoryName)
		newInventory := domain.NewInventoryResponse(invId, invName, newItemCategory, invQuantity, invPrice, invImage)

		inventory = append(inventory, newInventory)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration;", err)
	}

	return inventory, nil
}

func (this SQLiteInventoryRepository) GetById(id *types.InventoryId) (*domain.InventoryResponse, error) {
	stmt, err := this.db.Prepare(`
    SELECT i.id, i.name, i.item_category_id AS category_name, i.quantity, i.price, i.image ic.name AS category_name
    FROM inventory i 
    JOIN item_category ic ON i.item_category_id = ic.id
    WHERE i.id = ?;
  `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id.Value)

	var inv domain.InventoryPrimitive
	var itemCategory domainItemCategory.ItemCategoryPrimitive

	err = row.Scan(&inv.Id, &inv.Name, &itemCategory.Id, &inv.Quantity, &inv.Price, &inv.Image, &itemCategory.Name)
	if err != nil {
		return nil, err
	}

	invId, _ := types.NewInventoryId(inv.Id)
	invName, _ := types.NewInventoryName(inv.Name)
	invQuantity, _ := types.NewInventoryQuantity(inv.Quantity)
	invPrice, _ := types.NewInventoryPrice(inv.Price)
	invImage, _ := types.NewInventoryImage(inv.Image)

	itemCategoryId, _ := typesItemCategory.NewItemCategoryId(itemCategory.Id)
	itemCategoryName, _ := typesItemCategory.NewItemCategoryName(itemCategory.Name)

	newItemCategory := domainItemCategory.NewItemCategory(itemCategoryId, itemCategoryName)
	newInventory := domain.NewInventoryResponse(invId, invName, newItemCategory, invQuantity, invPrice, invImage)

	return newInventory, nil
}

func (this SQLiteInventoryRepository) Edit(inventory *domain.Inventory) error {
	stmt, err := this.db.Prepare("UPDATE inventory SET name = ?, item_category_id = ?, quantity = ?, price = ?, image = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(inventory.Name.Value, inventory.ItemCategoryId.Value, inventory.Quantity.Value, inventory.Price.Value, inventory.Image.Value, inventory.Id.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteInventoryRepository) Delete(id *types.InventoryId) error {
	stmt, err := this.db.Prepare("DELETE FROM inventory WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id.Value)
	if err != nil {
		return err
	}

	return nil
}
