package skuvariantsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/skuvariant"
	skuvariantrepo "fifentory/skuvariant/repository"
	"log"
)

const (
	skuVariantTable       = "sku_variant"
	createSKUVariantQuery = "INSERT " + skuVariantTable + " SET sku_id = ? , variant_value_id = ?"
)

func CreateSKUVariant(conn *sql.DB) skuvariantrepo.CreateSKUVariantFunc {
	return func(ctx context.Context, sv skuvariant.SKUVariant) (skuvariant.SKUVariant, error) {
		res, err := conn.ExecContext(ctx, createSKUVariantQuery, sv.SKUID, sv.VariantValueID)
		if err != nil {
			log.Println(err)
			return sv, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			return sv, err
		}
		sv.ID = id
		return sv, nil
	}
}
