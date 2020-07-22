package skupricesqlrepo

import (
	"context"
	"database/sql"
	"fifentory/skuprice"
	skupricerepo "fifentory/skuprice/repository"
	"log"
)

const (
	skuPriceTable       = "sku_price"
	createSKUPriceQuery = " INSERT INTO " + skuPriceTable + " SET sku_id = ? , selling_price = ? , buying_price"
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
