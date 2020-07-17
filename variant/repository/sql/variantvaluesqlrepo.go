package variantsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/variant"
	variantrepo "fifentory/variant/repository"
	"log"
)

const (
	variantValueTable  = "variant"
	variantValueFields = "id,value,value_option_id"
	getVariantsQuery   = "SELECT " + variantValueFields + " FROM " + variantValueTable
	createVariantQuery = "INSERT " + variantValueTable + " SET value = ? , variant_option_id = ?"
)

func CreateVariant(conn *sql.DB) variantrepo.CreateVariantFunc {
	return func(ctx context.Context, va variant.Variant) (variant.Variant, error) {
		res, err := conn.ExecContext(ctx, createVariantQuery, va.Value, va.Group.ID)
		if err != nil {
			log.Println(err)
			return va, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			return va, err
		}
		va.ID = id
		return va, nil
	}

}
