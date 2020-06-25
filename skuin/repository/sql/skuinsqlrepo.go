package skuinsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/skuin"
	skuinrepo "fifentory/skuin/repository"
)

const (
	skuInTable       = "sku_in"
	skuInFields      = "id,quantity,date,sku_id,sku_group_id"
	createSKUInQuery = "INSERT  " + skuInTable + " SET quantity = ? , date = ? , sku_id = ? , sku_in_group_id = ?"
)

func CreateSKUIn(conn *sql.DB) skuinrepo.CreateSKUInFunc {
	return func(ctx context.Context, in skuin.SKUIn) (skuin.SKUIn, error) {
		res, err := conn.ExecContext(ctx, createSKUInQuery, in.Quantity, in.Date, in.SKUID, in.GroupID)
		if err != nil {
			// error logging goes here
			return in, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return in, err
		}
		in.ID = id
		return in, nil
	}
}
