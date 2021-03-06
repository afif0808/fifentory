package skusqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/product"
	"fifentory/skuprice"
	"fifentory/stock"
	"fifentory/stockkeepingunit"
)

type receiver struct {
	SKU     *stockkeepingunit.SKU
	Product *product.Product
	Stock   *stock.Stock
	Price   *skuprice.SKUPrice
}

type SKUSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *receiver // used as receiver
	conn     *sql.DB
}

func (skuf *SKUSQLFetcher) Fetch(ctx context.Context, opts *options.Options) ([]stockkeepingunit.SKU, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	query := " SELECT " + skuf.fields + " FROM " + skuTable + " " + skuf.joins + " " + optionsQuery
	rows, err := skuf.conn.QueryContext(ctx, query, optionsArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	skus := []stockkeepingunit.SKU{}
	for rows.Next() {
		err := rows.Scan(skuf.scanDest...)
		if err != nil {
			return nil, err
		}
		sku := stockkeepingunit.SKU{
			ID:        skuf.Receiver.SKU.ID,
			Code:      skuf.Receiver.SKU.Code,
			CreatedAt: skuf.Receiver.SKU.CreatedAt,
		}
		if skuf.Receiver.Product != nil {
			pd := product.Product{
				ID:        skuf.Receiver.Product.ID,
				Name:      skuf.Receiver.Product.Name,
				CreatedAt: skuf.Receiver.Product.CreatedAt,
			}
			sku.Product = &pd
		}
		if skuf.Receiver.Stock != nil {
			st := stock.Stock{
				SKUID:           skuf.Receiver.Stock.SKUID,
				Quantity:        skuf.Receiver.Stock.Quantity,
				MinimumQuantity: skuf.Receiver.Stock.MinimumQuantity,
			}
			sku.Stock = &st
		}
		if skuf.Receiver.Price != nil {
			pr := skuprice.SKUPrice{
				SKUID:        skuf.Receiver.Price.SKUID,
				SellingPrice: skuf.Receiver.Price.SellingPrice,
				BuyingPrice:  skuf.Receiver.Price.BuyingPrice,
			}
			sku.Price = &pr
		}

		skus = append(skus, sku)
	}
	return skus, nil
}

func NewSKUSQLFetcher(conn *sql.DB) SKUSQLFetcher {
	sf := SKUSQLFetcher{
		joins:    "",
		fields:   "sku.id , sku.code , sku.created_at",
		Receiver: &receiver{SKU: &stockkeepingunit.SKU{}},
		conn:     conn,
	}
	sf.scanDest = []interface{}{&sf.Receiver.SKU.ID, &sf.Receiver.SKU.Code, &sf.Receiver.SKU.CreatedAt}
	return sf

}

func (skuf *SKUSQLFetcher) AddJoins(joins string) {
	skuf.joins += joins
}
func (skuf *SKUSQLFetcher) AddFields(fields string) {
	skuf.fields += fields
}
func (skuf *SKUSQLFetcher) AddScanDest(dest []interface{}) {
	skuf.scanDest = append(skuf.scanDest, dest...)
}
