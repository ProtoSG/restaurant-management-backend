package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"restaurant-management-backend/cmd/table_category/domain"
	"restaurant-management-backend/cmd/table_category/domain/types"
)

type SQLiteTableCategoryRepository struct {
	db *sql.DB
}

func NewSQLiteTableCategoryRepository(db *sql.DB) *SQLiteTableCategoryRepository {

	return &SQLiteTableCategoryRepository{db: db}
}

func (this SQLiteTableCategoryRepository) Create(tableCategory *domain.TableCategory) error {
	stmt, err := this.db.Prepare("INSERT INTO table_category (name) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tableCategory.Name.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteTableCategoryRepository) GetAll() ([]*domain.TableCategory, error) {
	stmt, err := this.db.Prepare("SELECT * FROM table_category")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tableCategories []*domain.TableCategory

	for rows.Next() {
		var tableCategory domain.TableCategoryPrimitive

		if err := rows.Scan(&tableCategory.Id, &tableCategory.Name); err != nil {
			return nil, err
		}

		tableCategoryId, _ := types.NewTableCategoryId(tableCategory.Id)
		tableCategoryName, _ := types.NewTableCategoryName(tableCategory.Name)

		newTableCategory := domain.NewTableCategory(tableCategoryId, tableCategoryName)

		tableCategories = append(tableCategories, newTableCategory)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration;", err)
	}

	return tableCategories, nil
}

func (this SQLiteTableCategoryRepository) GetById(tableCategoryId *types.TableCategoryId) (*domain.TableCategory, error) {
	stmt, err := this.db.Prepare("SELECT * FROM table_category WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(tableCategoryId.Value)

	var tableCategory domain.TableCategoryPrimitive

	err = row.Scan(&tableCategory.Id, &tableCategory.Name)
	if err != nil {
		return nil, err
	}

	tableCategoryName, _ := types.NewTableCategoryName(tableCategory.Name)

	newTableCategory := domain.NewTableCategory(tableCategoryId, tableCategoryName)

	return newTableCategory, nil
}

func (this SQLiteTableCategoryRepository) Edit(tableCategory *domain.TableCategory) error {
	stms, err := this.db.Prepare("UPDATE table_category SET name = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stms.Close()

	_, err = stms.Exec(tableCategory.Name.Value, tableCategory.Id.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteTableCategoryRepository) Delete(id *types.TableCategoryId) error {
	stmt, err := this.db.Prepare("DELETE FROM table_category WHERE id = ?")
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
