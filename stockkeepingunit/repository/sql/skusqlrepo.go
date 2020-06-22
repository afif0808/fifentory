package skusqlrepo

import (
	"context"
	"database/sql"
	"errors"
	"fifentory/options"
	"fifentory/stockkeepingunit"
	skurepo "fifentory/stockkeepingunit/repository"
)

const (
	skuTable           = "sku"
	skuFields          = "id,code,product_id,created_at"
	createSKUQuery     = "INSERT " + skuTable + " SET code = ? , product_id = ? ,created_at = ?"
	getSKUsQuery       = "SELECT " + skuFields + " FROM " + skuTable
	deleteSKUByIDQuery = "DELETE FROM sku WHERE id = ? "
	updateSKUByIDQuery = "UPDATE " + skuTable + " SET code = ? WHERE id = ?"
	getSKUByIDQuery    = "SELECT " + skuFields + " FROM " + skuTable + " WHERE id = ?"
)

// CreateSKU create a sku in sql database
// create product if caller
func CreateSKU(
	conn *sql.DB,
) skurepo.CreateSKUFunc {
	return func(ctx context.Context, sku stockkeepingunit.StockKeepingUnit) (stockkeepingunit.StockKeepingUnit, error) {
		//create sku
		res, err := conn.ExecContext(ctx, createSKUQuery, sku.Code, sku.ProductID, sku.CreatedAt)
		if err != nil {
			// error logging goes here
			return sku, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			// error logging goes here
			return sku, err
		}

		sku.ID = id

		return sku, nil
	}
}

func fetch(conn *sql.DB, ctx context.Context, query string, args ...interface{}) ([]stockkeepingunit.StockKeepingUnit, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	skus := []stockkeepingunit.StockKeepingUnit{}
	for rows.Next() {
		sku := stockkeepingunit.StockKeepingUnit{}
		err := rows.Scan(&sku.ID, &sku.Code, &sku.ProductID, &sku.CreatedAt)
		if err != nil {
			return nil, err
		}
		skus = append(skus, sku)
	}

	return skus, nil
}

func GetSKUs(
	conn *sql.DB,
) skurepo.GetSKUs {
	return func(ctx context.Context, opts *options.Options) ([]stockkeepingunit.StockKeepingUnit, error) {

		query, args := options.ParseOptionsToSQLQuery(opts)
		query = getSKUsQuery + query
		skus, err := fetch(conn, ctx, query, args...)
		if err != nil {
			// error logging goes here
			return nil, err
		}
		return skus, nil
	}
}
func DeleteSKUById(conn *sql.DB) skurepo.DeleteSKUByID {
	return func(ctx context.Context, id int64) error {
		_, err := conn.ExecContext(ctx, deleteSKUByIDQuery, id)
		if err != nil {
			// error logging goes here
			return err
		}
		return nil
	}
}
func UpdateSKUByID(conn *sql.DB) skurepo.UpdateSKUFunc {
	return func(ctx context.Context, sku stockkeepingunit.StockKeepingUnit) (stockkeepingunit.StockKeepingUnit, error) {
		_, err := conn.ExecContext(ctx, updateSKUByIDQuery, sku.Code, sku.ID)
		if err != nil {
			return (stockkeepingunit.StockKeepingUnit{}), err
		}
		return sku, nil
	}
}
func GetSKUByID(conn *sql.DB) skurepo.GetSKUByIDFunc {
	return func(ctx context.Context, id int64) (*stockkeepingunit.StockKeepingUnit, error) {
		skus, err := fetch(conn, ctx, getSKUByIDQuery, id)
		if err != nil {
			return nil, err
		}
		if len(skus) <= 0 {
			return nil, errors.New("sku with given id was not found")
		}

		return &skus[0], nil
	}
}
