package skuoutsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/skuout"
)

type receiver struct {
	SKUOut *skuout.SKUOut
}

type SKUOutSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *receiver
	conn     *sql.DB
}

func NewSKUOutSQLFetcher(conn *sql.DB) SKUOutSQLFetcher {
	sosf := SKUOutSQLFetcher{
		fields:   "sku_out.id,sku_out.quantity,sku_out.date",
		conn:     conn,
		Receiver: &receiver{SKUOut: &skuout.SKUOut{}},
	}
	sosf.scanDest = []interface{}{&sosf.Receiver.SKUOut.ID, &sosf.Receiver.SKUOut.Quantity, &sosf.Receiver.SKUOut.Date}
	return sosf
}

func (sosf *SKUOutSQLFetcher) Fetch(ctx context.Context, opts *options.Options) ([]skuout.SKUOut, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	rows, err := sosf.conn.QueryContext(ctx, " SELECT "+sosf.fields+" FROM "+skutOutTable+" "+sosf.joins+" "+optionsQuery, optionsArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sos := []skuout.SKUOut{}
	for rows.Next() {
		err = rows.Scan(sosf.scanDest...)
		if err != nil {
			return nil, err
		}
		so := skuout.SKUOut{
			ID:       sosf.Receiver.SKUOut.ID,
			Quantity: sosf.Receiver.SKUOut.Quantity,
			Date:     sosf.Receiver.SKUOut.Date,
		}
		sos = append(sos, so)
	}
	return sos, nil
}
