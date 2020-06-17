package categorysqlrepo

import (
	"context"
	"database/sql"
	"fifentory/category"
	categoryrepo "fifentory/category/repository"
)

const (
	categoryTable                      = "category"
	categoryFields                     = "id,value"
	getCategoryByIDQuery               = "SELECT " + categoryFields + " FROM " + categoryTable + " WHERE id = ?"
	productCategoryJunctionTable       = "product_category"
	productCategoryJunctionFields      = "id,product_id,category_id"
	getProductCategoryJunctionsQuery   = "SELECT " + productCategoryJunctionFields + " FROM " + productCategoryJunctionTable + " WHERE product_id = ?"
	createProductCategoryJunctionQuery = "INSERT " + productCategoryJunctionTable + " SET product_id = ?, category_id = ?"
)

// fetchProductCategoryJunctions fetch junction  of product and its category
func fetchProductCategoryJunctions(conn *sql.DB, ctx context.Context, query string, args ...interface{}) ([]category.ProductCategoryJunction, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	productCategoryJunctions := []category.ProductCategoryJunction{}
	for rows.Next() {
		pcj := category.ProductCategoryJunction{}
		err = rows.Scan(&pcj.ID, &pcj.ProductID, &pcj.CategoryID)
		if err != nil {
			return nil, err
		}
		productCategoryJunctions = append(productCategoryJunctions, pcj)
	}
	return productCategoryJunctions, nil
}

func fetchCategories(conn *sql.DB, ctx context.Context, query string, args ...interface{}) ([]category.Category, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	categories := []category.Category{}
	defer rows.Close()
	for rows.Next() {
		c := category.Category{}
		err := rows.Scan(&c.ID, &c.Value)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)

	}
	return categories, nil
}

func GetCategoryByID(conn *sql.DB) categoryrepo.GetCategoryByIDFunc {
	return categoryrepo.GetCategoryByIDFunc(func(ctx context.Context, id int64) (category.Category, error) {
		categories, err := fetchCategories(conn, ctx, getCategoryByIDQuery, id)
		if err != nil {
			//error logging goes here
			return (category.Category{}), err
		}
		if len(categories) <= 0 {
			return (category.Category{}), err
		}
		return categories[0], nil
	})
}

func GetProductCategories(
	conn *sql.DB,
	getCategory categoryrepo.GetCategoryByIDFunc,
) categoryrepo.GetProductCategoriesFunc {
	return categoryrepo.GetProductCategoriesFunc(func(ctx context.Context, productID int64) ([]category.Category, error) {
		productCategories := []category.Category{}

		junctions, err := fetchProductCategoryJunctions(conn, ctx, getProductCategoryJunctionsQuery, productID)
		if err != nil {

			// error logging goes here
			return nil, err
		}
		for _, v := range junctions {
			c, err := getCategory(ctx, v.CategoryID)
			if err != nil {
				// error logging goes here
				return nil, err
			}
			productCategories = append(productCategories, c)
		}
		return productCategories, nil
	})
}

func CreateProductCategoryJunction(
	conn *sql.DB,
) categoryrepo.CreateProductCategoryJunctionFunc {
	return func(ctx context.Context, junction category.ProductCategoryJunction) (category.ProductCategoryJunction, error) {
		res, err := conn.ExecContext(ctx, createProductCategoryJunctionQuery, junction.ProductID, junction.CategoryID)
		if err != nil {
			return junction, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return junction, err
		}
		junction.ID = id
		return junction, nil
	}
}
