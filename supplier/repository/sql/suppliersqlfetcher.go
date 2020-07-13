package suppliersqlrepo

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/supplier"
	"log"
)

// func SupplierSQLJoin(sf models.SQLFetcher, s *supplier.Supplier, foreignKey string) {
// 	dest := []interface{}{&s.ID, &s.Name}
// 	sf.AddScanDest(dest)
// 	sf.AddJoins(" INNER JOIN supplier ON supplier.id = " + foreignKey)
// 	sf.AddFields(",supplier.id , supplier.name")
// }
func SupplierSQLJoin(sf models.SQLFetcher, foreignKey string) *supplier.Supplier {
	s := &supplier.Supplier{}
	dest := []interface{}{&s.ID, &s.Name}
	sf.AddScanDest(dest)
	sf.AddJoins(" INNER JOIN supplier ON supplier.id = " + foreignKey)
	sf.AddFields(",supplier.id , supplier.name")
	return s
}

type SupplierSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *supplier.Supplier // used as receiver
	conn     *sql.DB
}

func (ssf *SupplierSQLFetcher) Fetch(ctx context.Context, opts *options.Options) ([]supplier.Supplier, error) {
	optionsQuery, optionsArg := options.ParseOptionsToSQLQuery(opts)
	query := "SELECT " + ssf.fields + " FROM supplier " + ssf.joins + " " + optionsQuery
	rows, err := ssf.conn.QueryContext(ctx, query, optionsArg...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	suppliers := []supplier.Supplier{}
	for rows.Next() {
		err := rows.Scan(ssf.scanDest...)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, *ssf.Receiver)
	}
	return suppliers, nil
}

func NewSupplierSQLFetcher(conn *sql.DB) SupplierSQLFetcher {
	ssf := SupplierSQLFetcher{
		joins:    "",
		fields:   "supplier.id , supplier.name",
		Receiver: &supplier.Supplier{},
		conn:     conn,
	}
	ssf.scanDest = []interface{}{&ssf.Receiver.ID, &ssf.Receiver.Name}
	return ssf
}

