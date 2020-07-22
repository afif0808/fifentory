package skuoutsqlrepo

import (
	"context"
	"database/sql"
	"errors"
	"fifentory/options"
	"fifentory/skuout"
	skuoutrepo "fifentory/skuout/repository"
	"log"
)

const (
	skutOutTable               = "sku_out"
	skuOutFields               = "id,quantity,date,sku_id,sku_out_group_id"
	createSKUOutQuery          = "INSERT " + skutOutTable + " SET quantity = ? , sku_id = ? ,date = ? , sku_out_group_id = ?"
	deleteSKUOutByIDQuery      = "DELETE FROM " + skutOutTable + " WHERE id = ?"
	getSKUOUtByIDQuery         = "SELECT " + skuOutFields + " FROM " + skutOutTable + " WHERE id = ? "
	getSKUOutsQuery            = "SELECT " + skuOutFields + " FROM " + skutOutTable
	getSKUOutsByGroupIdQuery   = "SELECT " + skuOutFields + " FROM " + skutOutTable + " WHERE sku_out_group_id = ?"
	updateSKUOutByIDQuery      = "UPDATE " + skutOutTable + " SET quantity = ? , date = ? , sku_id = ?, sku_out_group_id = ? " + " WHERE id = ? "
	deleteSKUOutBySKUIDQuery   = "DELETE FROM " + skutOutTable + " WHERE sku_id = ? "
	deleteSKUOutByGroupIDQuery = "DELETE FROM " + skutOutTable + " WHERE sku_out_group_id = ?"
)

func CreateSKUOut(conn *sql.DB) skuoutrepo.CreateSKUOutFunc {
	return func(ctx context.Context, out skuout.SKUOut) (skuout.SKUOut, error) {
		res, err := conn.ExecContext(ctx, createSKUOutQuery, out.Quantity, out.SKUID, out.Date, out.GroupID)
		if err != nil {
			return out, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return out, err
		}
		out.ID = id
		return out, nil
	}
}

func DeleteSKUOutByID(conn *sql.DB) skuoutrepo.DeleteSKUOutByIDFunc {
	return func(ctx context.Context, id int64) error {
		res, err := conn.ExecContext(ctx, deleteSKUOutByIDQuery, id)
		if err != nil {
			// error logging goes here
			return err
		}
		ra, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if ra > 1 {

		}
		return nil
	}
}

func fetchSKUOuts(conn *sql.DB, ctx context.Context, query string, args ...interface{}) ([]skuout.SKUOut, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { rows.Close() }()
	skuOuts := []skuout.SKUOut{}
	for rows.Next() {
		so := skuout.SKUOut{}
		err := rows.Scan(&so.ID, &so.Quantity, &so.Date, &so.SKUID, &so.GroupID)
		if err != nil {
			return nil, err
		}
		skuOuts = append(skuOuts, so)

	}
	return skuOuts, nil
}

func GetSKUOUtByID(conn *sql.DB) skuoutrepo.GetSKUOUtByIDFunc {
	return func(ctx context.Context, id int64) (*skuout.SKUOut, error) {
		skuOuts, err := fetchSKUOuts(conn, ctx, getSKUOUtByIDQuery, id)
		if err != nil {
			// error logging goes here
			return nil, err
		}
		if len(skuOuts) <= 0 {
			return nil, errors.New("Error : not found")
		}
		return &skuOuts[0], nil
	}
}
func GetSKUOuts(conn *sql.DB) skuoutrepo.GetSKUOutsFunc {
	return func(ctx context.Context, opts *options.Options) ([]skuout.SKUOut, error) {
		optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
		skuOuts, err := fetchSKUOuts(conn, ctx, getSKUOutsQuery+" "+optionsQuery, optionsArgs...)
		if err != nil {
			// error logging goes here
			return nil, err
		}
		return skuOuts, nil
	}
}
func GetSKUOutsByGroupId(conn *sql.DB) skuoutrepo.GetSKUOutsByGroupIDFunc {
	return func(ctx context.Context, groupID int64) ([]skuout.SKUOut, error) {
		skuOuts, err := fetchSKUOuts(conn, ctx, getSKUOutsByGroupIdQuery, groupID)
		if err != nil {
			return nil, err
		}
		return skuOuts, nil
	}
}
func UpdateSKUOutByID(conn *sql.DB) skuoutrepo.UpdateSKUOutFunc {
	return func(ctx context.Context, skuOut skuout.SKUOut) error {
		res, err := conn.ExecContext(ctx, updateSKUOutByIDQuery, skuOut.Quantity, skuOut.Date, skuOut.SKUID, skuOut.GroupID, skuOut.ID)
		if err != nil {
			return err
		}
		ra, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if ra != 1 {
			return errors.New("Error : something's wrong :D ")
		}
		return nil
	}
}
func DeleteSKUOutBySKUID(conn *sql.DB) skuoutrepo.DeleteSKUOutBySKUIDFunc {
	return func(ctx context.Context, skuID int64) error {
		_, err := conn.ExecContext(ctx, deleteSKUOutBySKUIDQuery, skuID)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}

func DeleteSKUOutByGroupID(conn *sql.DB) skuoutrepo.DeleteSKUOutByGroupIDFunc {
	return func(ctx context.Context, groupID int64) error {
		_, err := conn.ExecContext(ctx, deleteSKUOutByGroupIDQuery, groupID)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}
