package infrastructure

import (
	"database/sql"
	"fmt"
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/types"
	domainTableCategory "restaurant-management-backend/cmd/table_category/domain"
	typesTableCategory "restaurant-management-backend/cmd/table_category/domain/types"
)

type SQLiteTableRepository struct {
	db *sql.DB
}

func NewSQLiteTableRepository(db *sql.DB) *SQLiteTableRepository {
	return &SQLiteTableRepository{db: db}
}

func (this SQLiteTableRepository) Create(table *domain.Table) error {
	stmt, err := this.db.Prepare("INSERT INTO tables (name, table_category_id, status) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(table.Name.Value, table.CategoyId.Value, table.Status.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteTableRepository) GetAll() ([]*domain.TableResponse, error) {
	stmt, err := this.db.Prepare(`
      SELECT t.id, t.name, t.status, tc.id AS category_id, tc.name AS category_name
      FROM tables t 
      JOIN table_category tc ON t.table_category_id = tc.id;
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

	var tables []*domain.TableResponse

	for rows.Next() {
		var table domain.TablePrimitive
		var category domainTableCategory.TableCategoryPrimitive

		if err := rows.Scan(&table.Id, &table.Name, &table.Status, &category.Id, &category.Name); err != nil {
			return nil, err
		}

		tableId, _ := types.NewTableId(table.Id)
		tableName, _ := types.NewTableName(table.Name)
		tableStatus, _ := types.NewTableStatus(table.Status)

		categoryId, _ := typesTableCategory.NewTableCategoryId(category.Id)
		categoryName, _ := typesTableCategory.NewTableCategoryName(category.Name)

		newCategory := domainTableCategory.NewTableCategory(categoryId, categoryName)

		newTable := domain.NewTableResponse(tableId, tableName, newCategory, tableStatus)

		tables = append(tables, newTable)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration;", err)
	}

	return tables, nil
}

func (this SQLiteTableRepository) GetById(tableId *types.TableId) (*domain.TableResponse, error) {
	stmt, err := this.db.Prepare(`
      SELECT t.id, t.name, t.status, tc.id AS category_id, tc.name AS category_name
      FROM tables t 
      JOIN table_category tc ON t.table_category_id = tc.id
      WHERE t.id = ?;
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(tableId.Value)

	var table domain.TablePrimitive
	var category domainTableCategory.TableCategoryPrimitive

	err = row.Scan(&table.Id, &table.Name, &table.Status, &category.Id, &category.Name)
	if err != nil {
		return nil, err
	}

	tableName, _ := types.NewTableName(table.Name)
	tableStatus, _ := types.NewTableStatus(table.Status)

	categoryId, _ := typesTableCategory.NewTableCategoryId(category.Id)
	categoryName, _ := typesTableCategory.NewTableCategoryName(category.Name)

	newCategory := domainTableCategory.NewTableCategory(categoryId, categoryName)
	newTableResponse := domain.NewTableResponse(tableId, tableName, newCategory, tableStatus)

	return newTableResponse, nil
}

func (this SQLiteTableRepository) Edit(table *domain.Table) error {
	stms, err := this.db.Prepare("UPDATE tables SET name = ?, table_category_id = ?, status = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stms.Close()

	_, err = stms.Exec(table.Name.Value, table.CategoyId.Value, table.Status.Value, table.Id.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteTableRepository) Delete(id *types.TableId) error {
	stmt, err := this.db.Prepare("DELETE FROM tables WHERE id = ?")
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
