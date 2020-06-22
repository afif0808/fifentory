package productsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/product"
	productrepo "fifentory/product/repository"

	"github.com/xwb1989/sqlparser"
)

const (
	productTable           = "product"
	productFields          = "id,name,created_at"
	getProductByIdQuery    = "SELECT " + productFields + " FROM " + productTable + " WHERE id = ?"
	createProductQuery     = "INSERT " + productTable + " SET name = ? , created_at = ?"
	updateProductByIdQuery = "UPDATE " + productTable + " SET name  = ? WHERE id = ? "
)

func fetchProducts(conn *sql.DB, ctx context.Context, query string, args ...interface{}) ([]product.Product, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { rows.Close() }()
	products := []product.Product{}
	for rows.Next() {
		p := product.Product{}
		err = rows.Scan(&p.ID, &p.Name, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(
	conn *sql.DB,
) productrepo.GetProductByIDFunc {
	return productrepo.GetProductByIDFunc(func(ctx context.Context, id int64) (*product.Product, error) {
		if _, err := sqlparser.Parse(getProductByIdQuery); err != nil {
			return nil, err
		}
		products, err := fetchProducts(conn, ctx, getProductByIdQuery, id)
		if err != nil {
			return nil, err
			// error logging goes here
		}
		if len(products) <= 0 {
			return nil, models.ErrNotFound
			// error logging goes here
		}
		p := products[0]
		// defines categories resolver
		return &p, nil
	})
}

func CreateProduct(
	conn *sql.DB,
) productrepo.CreateProductFunc {
	return func(ctx context.Context, prod product.Product) (product.Product, error) {
		res, err := conn.ExecContext(ctx, createProductQuery, prod.Name, prod.CreatedAt)
		if err != nil {
			// error logging goes here
			return prod, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			// error logging goes here
			return prod, err
		}
		prod.ID = id
		return prod, nil
	}
}

func UpdateProductById(conn *sql.DB) productrepo.UpdateProductFunc {
	return func(ctx context.Context, prod product.Product) (product.Product, error) {
		_, err := conn.ExecContext(ctx, updateProductByIdQuery, prod.Name, prod.ID)
		if err != nil {
			return (product.Product{}), err
		}
		return prod, nil
	}
}

func GetProducts(conn *sql.DB) productrepo.GetProductsFunc {
	return func(ctx context.Context, opts *options.Options) ([]product.Product, error) {
		optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
		query := " SELECT " + productFields + " FROM " + productTable + " " + optionsQuery
		products, err := fetchProducts(conn, ctx, query, optionsArgs...)
		if err != nil {
			// logging goes here
			return nil, err
		}
		return products, nil
	}
}
