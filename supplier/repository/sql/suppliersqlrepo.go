package suppliersqlrepo

import (
	"context"
	"database/sql"
	"fifentory/supplier"
	supplierrepo "fifentory/supplier/repository"
)

const (
	supplierTable       = "supplier"
	createSupplierQuery = "INSERT " + supplierTable + " SET name = ?"
)

func CreateSupplier(conn *sql.DB) supplierrepo.CreateSupplierFunc {
	return func(ctx context.Context, sup supplier.Supplier) (supplier.Supplier, error) {
		res, err := conn.ExecContext(ctx, createSupplierQuery, sup.Name)
		if err != nil {
			return sup, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return sup, err
		}
		sup.ID = id
		return sup, nil
	}
}
