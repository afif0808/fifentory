package stocksqlrepo

import (
	"context"
	"database/sql"
	"fifentory/stock"
	stockrepo "fifentory/stock/repository"
)

const (
	skuStockFields               = "id,quantity"
	skuStockTable                = "sku_stock"
	createStockQuery             = "INSERT " + skuStockTable + " SET quantity = ? , sku_id = ?  "
	getSKUStockBySKUIDQuery      = "SELECT " + skuStockFields + " FROM " + skuStockTable + " WHERE sku_id = ?"
	updateSKUStockByIDQuery      = "UPDATE " + skuStockTable + " SET quantity = ? WHERE id = ? "
	addStockQuantityBySKUIDQuery = "UPDATE " + skuStockTable + " SET quantity = quantity + ? WHERE sku_id = ?"
)

func CreateStock(conn *sql.DB) stockrepo.CreateSKUStockFunc {
	return func(ctx context.Context, st stock.Stock) (stock.Stock, error) {
		res, err := conn.ExecContext(ctx, createStockQuery, st.Quantity, st.SKUID)
		if err != nil {
			return st, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return st, err
		}
		st.ID = id
		return st, nil
	}
}

func fetch(conn *sql.DB, ctx context.Context, query string, args ...interface{}) ([]stock.Stock, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	stocks := []stock.Stock{}
	for rows.Next() {
		st := stock.Stock{}
		err = rows.Scan(&st.ID, &st.Quantity)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, st)
	}
	return stocks, nil
}

func GetSKUStockBySKUID(
	conn *sql.DB,
) stockrepo.GetSKUStockBySKUID {
	return func(ctx context.Context, skuID int64) (*stock.Stock, error) {
		skus, err := fetch(conn, ctx, getSKUStockBySKUIDQuery, skuID)
		if err != nil {
			// error logging goes here
			return nil, err
		}
		if len(skus) <= 0 {
			return nil, err
		}
		return &skus[0], nil
	}
}
func UpdateSKUStockByID(
	conn *sql.DB,
) stockrepo.UpdateSKUStockFunc {
	return func(ctx context.Context, st stock.Stock) (stock.Stock, error) {
		_, err := conn.ExecContext(ctx, updateSKUStockByIDQuery, st.Quantity, st.ID)
		if err != nil {
			//error logging goes here
			return st, err
		}
		return st, nil
	}
}

func SubtractStockQuantityBySKUID(conn *sql.DB) stockrepo.SubtractStockQuantityBySKUIDFunc {
	return func(ctx context.Context, skuID int64, quantity int) error {
		_, err := conn.ExecContext(
			ctx, "UPDATE sku_stock SET quantity = quantity - ? WHERE sku_id  = ?",
			quantity, skuID)
		if err != nil {
			// error logging goes here
			return err
		}
		return nil
	}
}
func AddStockQuantityBySKUID(conn *sql.DB) stockrepo.AddStockQuantityBySKUIDFunc {
	return func(ctx context.Context, skuID int64, quantity int) error {
		res, err := conn.ExecContext(ctx, addStockQuantityBySKUIDQuery, quantity, skuID)
		if err != nil {
			// error logging goes here
			return err
		}
		ra, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if ra > 1 {
			// had no idea what to do
		}
		return nil
	}
}
