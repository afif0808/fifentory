package skuingroupsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/skuingroup"
	skuingrouprepo "fifentory/skuingroup/repository"
	"log"
)

const (
	skuInGroupTable        = "sku_in_group"
	createSKUInGroupQuery  = "INSERT " + skuInGroupTable + " SET supplier_id = ? , date = ? "
	deleteSKUInGroupsQuery = "DELETE FROM " + skuInGroupTable
)

func CreateSKUInGroup(conn *sql.DB) skuingrouprepo.CreateSKUInGroupFunc {
	return func(ctx context.Context, ingroup skuingroup.SKUInGroup) (skuingroup.SKUInGroup, error) {
		res, err := conn.ExecContext(ctx, createSKUInGroupQuery, ingroup.Supplier.ID, ingroup.Date)
		if err != nil {
			log.Println(err)
			return ingroup, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			return ingroup, err
		}
		ingroup.ID = id
		return ingroup, nil
	}
}
func DeleteSKUInGroups(conn *sql.DB) skuingrouprepo.DeleteSKUInGroupsFunc {
	return func(ctx context.Context, fts []options.Filter) error {
		filtersQuery, filtersArgs := options.ParseFiltersToSQLQuery(fts)
		query := deleteSKUInGroupsQuery + " " + filtersQuery
		_, err := conn.ExecContext(ctx, query, filtersArgs...)
		if err != nil {
			log.Println(err)
		}
		return err
	}
}
