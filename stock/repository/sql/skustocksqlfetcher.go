package stocksqlrepo

import (
	"database/sql"
	"fifentory/models"
	"fifentory/stock"
)

type receiver struct {
	SKUStock *stock.Stock
}

type SKUStockSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *receiver
	conn     *sql.DB
}

func NewSKUStockSQLFetcher(conn *sql.DB) SKUStockSQLFetcher {
	ssf := SKUStockSQLFetcher{
		fields:   "sku_stock.sku_id,sku_stock.quantity,sku_stock.minimum_quantity",
		conn:     conn,
		Receiver: &receiver{SKUStock: &stock.Stock{}},
	}
	ssf.scanDest = []interface{}{ssf.Receiver.SKUStock.SKUID, ssf.Receiver.SKUStock.Quantity, ssf.Receiver.SKUStock.MinimumQuantity}
	return ssf
}

func SKUStockSQLJoin(sf models.SQLFetcher) *stock.Stock {
	s := &stock.Stock{}
	dest := []interface{}{&s.SKUID, &s.Quantity, &s.MinimumQuantity}
	sf.AddScanDest(dest)
	sf.AddJoins(" INNER JOIN sku_stock ON sku_stock.sku_id = sku.id")
	sf.AddFields(",sku_stock.sku_id , sku_stock.quantity , sku_stock.minimum_quantity")
	return s
}
