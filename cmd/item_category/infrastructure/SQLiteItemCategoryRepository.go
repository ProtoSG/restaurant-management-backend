package infrastructure

import (
	"database/sql"
	"restaurant-management-backend/cmd/item_category/domain"
	"restaurant-management-backend/cmd/item_category/domain/types"
)

type SQLiteItemCategoryRepository struct {
	db *sql.DB
}

func NewSQLiteItemCategoryRepository(db *sql.DB) *SQLiteItemCategoryRepository {
	return &SQLiteItemCategoryRepository{db: db}
}

func (this *SQLiteItemCategoryRepository) Create(itemCategory *domain.ItemCategory) error {
	stmt, err := this.db.Prepare("INSERT INTO item_category (name) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(itemCategory.Name.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteItemCategoryRepository) GetAll() ([]*domain.ItemCategory, error) {
	stmt, err := this.db.Prepare("SELECT * FROM item_category")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itemCategories []*domain.ItemCategory

	for rows.Next() {
		var itemCategory domain.ItemCategoryPrimitive
		if err := rows.Scan(&itemCategory.Id, &itemCategory.Name); err != nil {
			return nil, err
		}

		itemCategoryId, _ := types.NewItemCategoryId(itemCategory.Id)
		itemCategoryName, _ := types.NewItemCategoryName(itemCategory.Name)

		newItemCategory := domain.NewItemCategory(itemCategoryId, itemCategoryName)

		itemCategories = append(itemCategories, newItemCategory)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return itemCategories, nil
}

func (this *SQLiteItemCategoryRepository) GetById(id *types.ItemCategoryId) (*domain.ItemCategory, error) {
	stmt, err := this.db.Prepare("SELECT * FROM item_category WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id.Value)

	var itemCategory domain.ItemCategoryPrimitive

	err = row.Scan(&itemCategory.Id, &itemCategory.Name)
	if err != nil {
		return nil, err
	}

	itemCategoryId, _ := types.NewItemCategoryId(itemCategory.Id)
	itemCategoryName, _ := types.NewItemCategoryName(itemCategory.Name)

	newItemCategory := domain.NewItemCategory(itemCategoryId, itemCategoryName)

	return newItemCategory, nil
}

func (this *SQLiteItemCategoryRepository) Edit(itemCategory *domain.ItemCategory) error {
	stmt, err := this.db.Prepare("UPDATE item_category SET name = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(itemCategory.Name.Value, itemCategory.Id.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteItemCategoryRepository) Delete(id *types.ItemCategoryId) error {
	stmt, err := this.db.Prepare("DELETE FROM item_category WHERE id = ?")
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
