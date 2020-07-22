package skupricesqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/skuprice"
	skupricerepo "fifentory/skuprice/repository"
	"log"
)

const (
	skuPriceTable       = "sku_price"
	createSKUPriceQuery = " INSERT INTO " + skuPriceTable + " SET sku_id = ? , selling_price = ? , buying_price = ?"
)

func CreateSKUPrice(conn *sql.DB) skupricerepo.CrateSKUPriceFunc {
	return func(ctx context.Context, sp skuprice.SKUPrice) (skuprice.SKUPrice, error) {
		_, err := conn.ExecContext(ctx, createSKUPriceQuery, sp.SKUID, sp.SellingPrice, sp.BuyingPrice)
		if err != nil {
			log.Println(err)
			return sp, err
		}
		return sp, nil
	}
}

func UpdateSKUPrice(conn *sql.DB) skupricerepo.UpdateSKUPriceFunc {
	return func(ctx context.Context, sp skuprice.SKUPrice, fts []options.Filter) (skuprice.SKUPrice, error) {
		filtersQuery, filtersArgs := options.ParseFiltersToSQLQuery(fts)
		filtersArgs = append([]interface{}{sp.SellingPrice, sp.BuyingPrice}, filtersArgs...)

		_, err := conn.ExecContext(
			ctx,
			"UPDATE "+skuPriceTable+" SET selling_price = ? , buying_price = ? "+filtersQuery,
			filtersArgs...,
		)
		if err != nil {
			log.Println(err)
			return sp, err
		}
		return sp, nil
	}
}
