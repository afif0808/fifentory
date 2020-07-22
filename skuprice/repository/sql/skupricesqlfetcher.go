package skupricesqlrepo

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/skuprice"
)

type receiver struct {
	SKUPrice *skuprice.SKUPrice
}

type SKUPriceSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *receiver
	conn     *sql.DB
}

func NewSKUPriceSQLFetcher(conn *sql.DB) SKUPriceSQLFetcher {
	spsf := SKUPriceSQLFetcher{
		fields:   "sku_price.sku_id,sku_price.selling_price,sku_price.buying_price",
		conn:     conn,
		Receiver: &receiver{SKUPrice: &skuprice.SKUPrice{}},
	}
	spsf.scanDest = []interface{}{
		&spsf.Receiver.SKUPrice.SKUID,
		&spsf.Receiver.SKUPrice.SellingPrice,
		&spsf.Receiver.SKUPrice.BuyingPrice,
	}
	return spsf
}

func (spsf *SKUPriceSQLFetcher) AddJoins(joins string) {
	spsf.joins += joins
}
func (spsf *SKUPriceSQLFetcher) AddFields(fields string) {
	spsf.fields += fields
}
func (spsf *SKUPriceSQLFetcher) AddScanDest(dest []interface{}) {
	spsf.scanDest = append(spsf.scanDest, dest...)
}

func (spsf *SKUPriceSQLFetcher) Fetch(ctx context.Context, opts *options.Options) ([]skuprice.SKUPrice, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	rows, err := spsf.conn.QueryContext(ctx, "SELECT "+spsf.fields+" FROM "+spsf.joins+" "+optionsQuery, optionsArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sps := []skuprice.SKUPrice{}
	for rows.Next() {
		err = rows.Scan(spsf.scanDest...)
		if err != nil {
			return nil, err
		}
		sp := skuprice.SKUPrice{
			SKUID:        spsf.Receiver.SKUPrice.SKUID,
			SellingPrice: spsf.Receiver.SKUPrice.SellingPrice,
			BuyingPrice:  spsf.Receiver.SKUPrice.BuyingPrice,
		}
		sps = append(sps, sp)
	}
	return sps, nil
}
func SKUPriceSQLJoin(sf models.SQLFetcher) *skuprice.SKUPrice {
	sp := &skuprice.SKUPrice{}
	dest := []interface{}{&sp.SKUID, &sp.SellingPrice, &sp.BuyingPrice}
	sf.AddScanDest(dest)
	sf.AddJoins(" INNER JOIN sku_price ON sku_price.sku_id = sku.id ")
	sf.AddFields(",sku_price.sku_id,sku_price.selling_price,sku_price.buying_price")
	return sp
}
