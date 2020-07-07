package skuinsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/skuin"
	skuinrepo "fifentory/skuin/repository"
	"log"
)

const (
	skuInTable       = "sku_in"
	skuInFields      = "id,quantity,date,sku_id,sku_group_id"
	createSKUInQuery = "INSERT  " + skuInTable + " SET quantity = ? , date = ? , sku_id = ? , sku_in_group_id = ?"
	deletSKUInQuery  = " DELETE FROM " + skuInTable
	updateSKUInQuery = "UPDATE " + skuInTable + " SET quantity = ? "
)

func CreateSKUIn(conn *sql.DB) skuinrepo.CreateSKUInFunc {
	return func(ctx context.Context, in skuin.SKUIn) (skuin.SKUIn, error) {
		res, err := conn.ExecContext(ctx, createSKUInQuery, in.Quantity, in.Date, in.SKU.ID, in.Group.ID)
		if err != nil {
			// error logging goes here
			log.Println(err)
			return in, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Println(err)

			return in, err
		}
		in.ID = id
		return in, nil
	}
}

func DeleteSKUIn(conn *sql.DB) skuinrepo.DeleteSKUInFunc {
	return func(ctx context.Context, fts []options.Filter) error {
		filtersQuery, filtersArgs := options.ParseFiltersToSQLQuery(fts)

		query := deletSKUInQuery + " " + filtersQuery
		_, err := conn.ExecContext(ctx, query, filtersArgs...)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}

func UpdateSKUIn(conn *sql.DB) skuinrepo.UpdateSKUInFunc {
	return func(ctx context.Context, in skuin.SKUIn, fts []options.Filter) (skuin.SKUIn, error) {
		filtersQuery, filtersArgs := options.ParseFiltersToSQLQuery(fts)
		query := updateSKUInQuery + " " + filtersQuery
		filtersArgs = append([]interface{}{in.Quantity}, filtersArgs...)
		_, err := conn.ExecContext(ctx, query, filtersArgs...)
		if err != nil {
			log.Println(err)
			return in, err
		}
		return in, nil
	}
}
