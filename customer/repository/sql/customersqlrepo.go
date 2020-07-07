package customersqlrepo

import (
	"context"
	"database/sql"
	"fifentory/customer"
	customerrepo "fifentory/customer/repository"
	"fifentory/models"
	"log"
)

const (
	customerTable        = "customer"
	customerFields       = "id,name"
	getCustomersQuery    = "SELECT " + customerFields + " FROM " + customerTable
	getCustomerByIDQuery = "SELECT " + customerFields + " FROM " + customerTable + " WHERE id = ?"
)

func fetchCustomers(conn *sql.DB, ctx context.Context, query string, args ...interface{}) ([]customer.Customer, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	customers := []customer.Customer{}
	for rows.Next() {
		c := customer.Customer{}
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
func GetCustomers(conn *sql.DB) customerrepo.GetCustomersFunc {
	return func(ctx context.Context) ([]customer.Customer, error) {
		customers, err := fetchCustomers(conn, ctx, getCustomersQuery)
		if err != nil {
			return nil, err
		}
		return customers, nil
	}
}

func GetCustomer(conn *sql.DB) customerrepo.GetCustomerByIDFunc {
	return func(ctx context.Context, id int64) (*customer.Customer, error) {
		customers, err := fetchCustomers(conn, ctx, getCustomerByIDQuery, id)
		if err != nil {
			return nil, err
		}
		if len(customers) <= 0 {
			return nil, nil
		}
		return &customers[0], nil
	}
}

func CustomerSQLJoin(sf models.SQLFetcher, c *customer.Customer, foreignKey string) {
	dest := []interface{}{&c.ID, &c.Name}
	sf.AddScanDest(dest)
	sf.AddJoins(" INNER JOIN customer ON customer.id = " + foreignKey)
	sf.AddFields(",customer.id , customer.name")
}
