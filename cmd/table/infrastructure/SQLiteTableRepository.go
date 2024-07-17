package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"os"
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/types"
)

type SQLiteTableRepository struct {
	db *sql.DB
}

func NewSQLiteTableRepository(url string) *SQLiteTableRepository {
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	return &SQLiteTableRepository{db: db}
}

func (this SQLiteTableRepository) Create(table *domain.Table) error {
	stmt, err := this.db.Prepare("INSERT INTO tables (name, table_category_id, status) VALUES (?. ?, ?)")
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

func (this SQLiteTableRepository) GetAll() ([]*domain.Table, error) {
	stmt, err := this.db.Prepare("SELECT * FROM tables")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []*domain.Table

	for rows.Next() {
		var table domain.TablePrimitive

		if err := rows.Scan(&table.Id, &table.Name, &table.CategoryId, &table.Status); err != nil {
			return nil, err
		}

		tableId, _ := types.NewTableId(table.Id)
		tableName, _ := types.NewTableName(table.Name)
		tableCategoryId, _ := types.NewTableCategoryId(table.CategoryId)
		tableStatus, _ := types.NewTableStatus(table.Status)

		newTable := domain.NewTable(tableId, tableName, tableCategoryId, tableStatus)

		tables = append(tables, newTable)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration;", err)
	}

	return tables, nil
}

func (this SQLiteTableRepository) GetById(tableId *types.TableId) (*domain.TableResponse, error) {
	stmt, err := this.db.Prepare("SELECT * FROM tables WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(tableId.Value)

	var table domain.TablePrimitive

	err = row.Scan(&table.Id, &table.Name, &table.CategoryId, &table.Status)
	if err != nil {
		return nil, err
	}

	tableName, _ := types.NewTableName(table.Name)
	tableCategoryId, _ := types.NewTableCategoryId(table.CategoryId)
	tableStatus, _ := types.NewTableStatus(table.Status)

	newCategory := types.NewTableCategory(tableCategoryId.Value, table.Name)

	newTable := domain.NewTableResponse(tableId, tableName, newCategory, tableStatus)

	return newTable, nil
}

func (this SQLiteTableRepository) Edit(table *domain.Table) error {
	stms, err := this.db.Prepare("UPDATE tables SET name = ?, table_category_id = ?, status, = ? WHERE id = ?")
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
