package productsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/product"
)

type receiver struct {
	Product *product.Product
}

type ProductSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *receiver
	conn     *sql.DB
}

func NewProductSQLFetcher(conn *sql.DB) ProductSQLFetcher {
	psf := ProductSQLFetcher{
		fields:   "product.id , product.name , product.created_at",
		conn:     conn,
		Receiver: &receiver{Product: &product.Product{}},
	}
	psf.scanDest = []interface{}{
		&psf.Receiver.Product.ID,
		&psf.Receiver.Product.Name,
		&psf.Receiver.Product.CreatedAt,
	}
	return psf
}

func (psf *ProductSQLFetcher) Fetch(ctx context.Context, opts *options.Options) ([]product.Product, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	rows, err := psf.conn.QueryContext(ctx, "SELECT "+psf.fields+" "+productTable+" "+psf.joins+" "+optionsQuery, optionsArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ps := []product.Product{}
	for rows.Next() {
		err = rows.Scan(psf.scanDest...)
		if err != nil {
			return nil, err
		}
		p := product.Product{
			ID:        psf.Receiver.Product.ID,
			Name:      psf.Receiver.Product.Name,
			CreatedAt: psf.Receiver.Product.CreatedAt,
		}
		ps = append(ps, p)
	}
	return ps, nil
}

func (psf *ProductSQLFetcher) AddJoins(joins string) {
	psf.joins += joins
}
func (psf *ProductSQLFetcher) AddFields(fields string) {
	psf.fields += fields
}
func (psf *ProductSQLFetcher) AddScanDest(dest []interface{}) {
	psf.scanDest = append(psf.scanDest, dest...)
}

func ProductSQLJoin(sf models.SQLFetcher, foreignKey string) *product.Product {
	p := &product.Product{}
	dest := []interface{}{&p.ID, &p.Name, &p.CreatedAt}
	sf.AddScanDest(dest)
	sf.AddJoins(" INNER JOIN product ON product.id = " + foreignKey)
	sf.AddFields(",product.id,product.name,product.created_at")
	return p
}
