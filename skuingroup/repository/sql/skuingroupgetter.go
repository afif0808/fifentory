package skuingroupsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/skuingroup"
	"fifentory/supplier"
	"log"
)

type SKUInGroupFetcher struct {
	fields     string
	joins      string
	scanDest   []interface{}
	SKUInGroup *skuingroup.SKUInGroup
	conn       *sql.DB
}

func NewSKUInGroupFetcher(conn *sql.DB) SKUInGroupFetcher {
	skuigf := SKUInGroupFetcher{
		conn:       conn,
		SKUInGroup: &skuingroup.SKUInGroup{},
		fields:     "sku_in_group.id,sku_in_group.date",
	}
	skuigf.scanDest = []interface{}{&skuigf.SKUInGroup.ID, &skuigf.SKUInGroup.Date}
	return skuigf
}
func (skuigf *SKUInGroupFetcher) Fetch(ctx context.Context, opts *options.Options) ([]skuingroup.SKUInGroup, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	query := "SELECT " + skuigf.fields + " FROM sku_in_group " + skuigf.joins + " " + optionsQuery
	rows, err := skuigf.conn.QueryContext(ctx, query, optionsArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	skuInGroups := []skuingroup.SKUInGroup{}
	for rows.Next() {
		err = rows.Scan(skuigf.scanDest...)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		skuInGroup := skuingroup.SKUInGroup{
			ID:   skuigf.SKUInGroup.ID,
			Date: skuigf.SKUInGroup.Date,
		}
		if skuigf.SKUInGroup.Supplier != nil {
			sp := supplier.Supplier{
				ID:   skuigf.SKUInGroup.Supplier.ID,
				Name: skuigf.SKUInGroup.Supplier.Name,
			}
			skuInGroup.Supplier = &sp
		}
		skuInGroups = append(skuInGroups, skuInGroup)
	}
	return skuInGroups, nil
}
func (skuigf *SKUInGroupFetcher) AddJoins(joins string) {
	skuigf.joins += joins
}
func (skuigf *SKUInGroupFetcher) AddFields(fields string) {
	skuigf.fields += fields
}
func (skuigf *SKUInGroupFetcher) AddScanDest(dest []interface{}) {
	skuigf.scanDest = append(skuigf.scanDest, dest...)
}
