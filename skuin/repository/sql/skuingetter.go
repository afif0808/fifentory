package skuinsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/product"
	"fifentory/skuin"
	"fifentory/skuingroup"
	"fifentory/stockkeepingunit"
)

type SKUInGetter struct {
	fields          string
	joins           string
	scanDestination []interface{}
	conn            *sql.DB
	*skuin.SKUIn
}

func (skuig *SKUInGetter) Get(ctx context.Context, opts *options.Options) ([]skuin.SKUIn, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	query := "SELECT " + skuig.fields + " FROM sku_in " + skuig.joins + optionsQuery
	rows, err := skuig.conn.QueryContext(ctx, query, optionsArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	skuIns := []skuin.SKUIn{}
	for rows.Next() {
		err = rows.Scan(skuig.scanDestination...)
		if err != nil {
			return nil, err
		}
		skuIn := skuin.SKUIn{
			ID:       skuig.SKUIn.ID,
			Date:     skuig.SKUIn.Date,
			Quantity: skuig.SKUIn.Quantity,
		}
		if skuig.SKU != nil {
			skuIn.SKU = &stockkeepingunit.SKU{
				ID:        skuig.SKU.ID,
				Code:      skuig.SKU.Code,
				CreatedAt: skuig.SKU.CreatedAt,
			}
			if skuig.SKU.Product != nil {
				skuIn.SKU.Product = &product.Product{
					ID:        skuig.SKU.Product.ID,
					Name:      skuig.SKU.Product.Name,
					CreatedAt: skuig.SKU.Product.CreatedAt,
				}
			}
		}

		if skuig.Group != nil {
			skuIn.Group = &skuingroup.SKUInGroup{
				ID:   skuig.Group.ID,
				Date: skuig.Group.Date,
			}
		}

		skuIns = append(skuIns, skuIn)
	}
	return skuIns, nil
}

func (skuig *SKUInGetter) WithSKU() {
	skuig.fields += ",sku.id,sku.code,sku.created_at"
	skuig.joins += " INNER JOIN sku ON sku.id = sku_in.sku_id"
	skuig.SKU = &stockkeepingunit.SKU{}
	sd := []interface{}{&skuig.SKUIn.SKU.ID, &skuig.SKU.Code, &skuig.SKU.CreatedAt}
	skuig.scanDestination = append(skuig.scanDestination, sd...)
}
func (skuig *SKUInGetter) WithProduct() {
	skuig.fields += ",product.id , product.name , product.created_at"
	skuig.joins += " INNER JOIN product ON product.id = sku.product_id"
	skuig.SKU.Product = &product.Product{}
	sd := []interface{}{&skuig.SKUIn.SKU.Product.ID, &skuig.SKUIn.SKU.Product.Name, &skuig.SKUIn.SKU.Product.CreatedAt}
	skuig.scanDestination = append(skuig.scanDestination, sd...)
}
func (skuig *SKUInGetter) WithGroup() {
	skuig.fields += ",sku_in.sku_in_group_id ,sku_in_group.id , sku_in_group.date"
	skuig.joins += " INNER JOIN sku_in_group ON sku_in.sku_in_group_id = sku_in_group.id"
	skuig.Group = &skuingroup.SKUInGroup{}
	sd := []interface{}{&skuig.SKUIn.ID, &skuig.SKUIn.Group.ID, &skuig.SKUIn.Group.Date}
	skuig.scanDestination = append(skuig.scanDestination, sd...)
}

func NewSKUInGetter(conn *sql.DB) SKUInGetter {
	skuig := SKUInGetter{
		fields: "sku_in.id , sku_in.quantity , sku_in.date",
		conn:   conn,
		SKUIn:  &skuin.SKUIn{},
	}
	skuig.scanDestination = []interface{}{&skuig.SKUIn.ID, &skuig.SKUIn.Quantity, &skuig.SKUIn.Date}
	return skuig
}
