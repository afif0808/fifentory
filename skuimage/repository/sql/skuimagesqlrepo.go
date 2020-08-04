package skuimagesqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/skuimage"
	skuimagerepo "fifentory/skuimage/repository"
	"log"
)

const (
	skuImageTable        = "sku_image"
	createSKUImageQuery  = " INSERT " + skuImageTable + " SET sku_id = ? , path = ?"
	deleteSKUImagesQuery = "DELETE FROM " + skuImageTable + " "
)

func CreateSKUImage(conn *sql.DB) skuimagerepo.CreateSKUImageFunc {
	return func(ctx context.Context, sim skuimage.SKUImage) (skuimage.SKUImage, error) {
		res, err := conn.ExecContext(ctx, createSKUImageQuery, sim.SKUID, sim.Path)
		if err != nil {
			log.Println(err)
			return sim, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			return sim, err
		}
		sim.ID = id
		return sim, nil
	}
}

func DeleteSKUImages(conn *sql.DB) skuimagerepo.DeleteSKUImagesFunc {
	return func(ctx context.Context, fts []options.Filter) error {
		filtersQuery, filtersArgs := options.ParseFiltersToSQLQuery(fts)
		_, err := conn.ExecContext(ctx, deleteSKUImagesQuery+filtersQuery, filtersArgs...)
		if err != nil {
			log.Println(err)
		}
		return err
	}
}
