package skuoutgroupsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/skuoutgroup"
	skuoutgrouprepo "fifentory/skuoutgroup/repository"
	"log"
)

const (
	skuOutGroupTable           = "sku_out_group"
	skuOutGroupFields          = "id,date,customer_id"
	createSKUOutGroupQuery     = "INSERT " + skuOutGroupTable + " SET date = ? , customer_id = ?"
	getSKUOutGroupsQuery       = "SELECT " + skuOutGroupFields + " FROM " + skuOutGroupTable
	deleteSKUOutGroupByIDQuery = "DELETE FROM " + skuOutGroupTable + " WHERE id = ?"
)

func CreateSKUOutGroup(conn *sql.DB) skuoutgrouprepo.CreateSKUOutGroupFunc {
	return func(ctx context.Context, outgroup skuoutgroup.SKUOutGroup) (skuoutgroup.SKUOutGroup, error) {
		res, err := conn.ExecContext(ctx, createSKUOutGroupQuery, outgroup.Date, outgroup.CustomerID)
		if err != nil {
			return outgroup, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return outgroup, err
		}
		outgroup.ID = id
		return outgroup, nil
	}
}
func fetchSKUOutGroup(conn *sql.DB, ctx context.Context, query string, args ...interface{}) ([]skuoutgroup.SKUOutGroup, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
	}()
	outgroups := []skuoutgroup.SKUOutGroup{}
	for rows.Next() {
		og := skuoutgroup.SKUOutGroup{}
		err := rows.Scan(&og.ID, &og.Date, &og.CustomerID)
		if err != nil {
			return nil, err
		}
		outgroups = append(outgroups, og)
	}

	return outgroups, nil
}
func GetSKUOutGroups(conn *sql.DB) skuoutgrouprepo.GetSKUOutGroupsFunc {
	return func(ctx context.Context) ([]skuoutgroup.SKUOutGroup, error) {
		outgroups, err := fetchSKUOutGroup(conn, ctx, getSKUOutGroupsQuery)
		if err != nil {
			return nil, err
		}
		return outgroups, nil
	}
}

func DeleteSKUOutGroupByID(conn *sql.DB) skuoutgrouprepo.DeleteSKUOutGroupByIDFunc {
	return func(ctx context.Context, id int64) error {
		_, err := conn.ExecContext(ctx, deleteSKUOutGroupByIDQuery, id)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}
