package stocksqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/stock"
	stockrepo "fifentory/stock/repository"
	"log"
)

const (
	skuStockFields               = "quantity,minimum_quantity,sku_id"
	skuStockTable                = "sku_stock"
	createStockQuery             = "INSERT " + skuStockTable + " SET quantity = ? , sku_id = ? , minimum_quantity = ? "
	getSKUStockBySKUIDQuery      = "SELECT " + skuStockFields + " FROM " + skuStockTable + " WHERE sku_id = ?"
	updateSKUStockByIDQuery      = "UPDATE " + skuStockTable + " SET quantity = ? , minimum_quantity = ? WHERE sku_id = ?  "
	addStockQuantityBySKUIDQuery = "UPDATE " + skuStockTable + " SET quantity = quantity + ? WHERE sku_id = ?"
	getRunningLowStocksQuery     = "SELECT " + skuStockFields + " FROM " + skuStockTable + " WHERE quantity <= minimum_quantity "
	deleteStockBySKUIDQuery      = "DELETE FROM " + skuStockTable + " WHERE sku_id = ?"
	substractStockQuantityQuery  = "UPDATE " + skuStockTable + " SET quantity = quantity - ?"
	addStockQuantityQuery        = "UPDATE " + skuStockTable + " SET quantity = quantity + ?"
)

func CreateStock(conn *sql.DB) stockrepo.CreateSKUStockFunc {
	return func(ctx context.Context, st stock.Stock) (stock.Stock, error) {
		_, err := conn.ExecContext(ctx, createStockQuery, st.Quantity, st.SKUID, st.MinimumQuantity)
		if err != nil {
			return st, err
		}
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
		err = rows.Scan(&st.Quantity, &st.MinimumQuantity, &st.SKUID)
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
func UpdateSKUStockBySKUID(
	conn *sql.DB,
) stockrepo.UpdateSKUStockFunc {
	return func(ctx context.Context, st stock.Stock) (stock.Stock, error) {
		_, err := conn.ExecContext(ctx, updateSKUStockByIDQuery, st.Quantity, st.MinimumQuantity, st.SKUID)
		if err != nil {
			//error logging goes here
			log.Println(err)
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

func GetRunningLowStocks(conn *sql.DB) stockrepo.GetRunnigLowStocksFunc {
	return func(ctx context.Context) ([]stock.Stock, error) {
		stocks, err := fetch(conn, ctx, getRunningLowStocksQuery)
		if err != nil {
			return nil, err
		}
		return stocks, nil
	}
}

func DeleteStockBySKUID(conn *sql.DB) stockrepo.DeleteStockBySKUID {
	return func(ctx context.Context, skuID int64) error {
		_, err := conn.ExecContext(ctx, deleteStockBySKUIDQuery, skuID)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}

func SubstractStockQuantity(conn *sql.DB) stockrepo.SubtractStockQuantityFunc {
	return func(ctx context.Context, fts []options.Filter, quantity int) error {
		filtersQuery, filtersArgs := options.ParseFiltersToSQLQuery(fts)
		filtersArgs = append([]interface{}{quantity}, filtersArgs...)
		query := substractStockQuantityQuery + " " + filtersQuery
		_, err := conn.ExecContext(ctx, query, filtersArgs...)
		if err != nil {
			log.Println(err)
		}
		return err
	}
}

func AddStockQuantity(conn *sql.DB) stockrepo.AddStockQuantityFunc {
	return func(ctx context.Context, fts []options.Filter, quantity int) error {
		filtersQuery, filtersArgs := options.ParseFiltersToSQLQuery(fts)
		filtersArgs = append([]interface{}{quantity}, filtersArgs...)
		query := addStockQuantityQuery + " " + filtersQuery
		_, err := conn.ExecContext(ctx, query, filtersArgs...)
		if err != nil {
			log.Println(err)
		}
		return err
	}
}
