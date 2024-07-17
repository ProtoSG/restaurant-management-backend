package domain

import (
	"fmt"
	"restaurant-management-backend/cmd/table/domain/types"
)

type TableNotFound struct {
	TableId types.TableId
}

func NewTableNotFound(tableId types.TableId) *TableNotFound {
	return &TableNotFound{TableId: tableId}
}

func (this TableNotFound) Error() string {
	return fmt.Sprintf("Tabla con el ID %d no existe", this.TableId)
}
