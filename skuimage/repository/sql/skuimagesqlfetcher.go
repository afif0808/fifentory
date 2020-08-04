package skuimagesqlrepo

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/skuimage"
)

type receiver struct {
	SKUImage *skuimage.SKUImage
}
type SKUImageSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *receiver // used as receiver
	conn     *sql.DB
}

func NewSKUImageSQLFetcher(conn *sql.DB) SKUImageSQLFetcher {
	sims := SKUImageSQLFetcher{
		fields:   "sku_image.id , sku_image.path , sku_image.sku_id",
		conn:     conn,
		Receiver: &receiver{SKUImage: &skuimage.SKUImage{}},
	}
	sims.scanDest = []interface{}{
		&sims.Receiver.SKUImage.ID,
		&sims.Receiver.SKUImage.Path,
		&sims.Receiver.SKUImage.SKUID,
	}
	return sims
}

func (simsf *SKUImageSQLFetcher) Fetch(ctx context.Context, opts *options.Options) ([]skuimage.SKUImage, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	rows, err := simsf.conn.QueryContext(ctx, " SELECT "+simsf.fields+" FROM "+skuImageTable+" "+optionsQuery, optionsArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sims := []skuimage.SKUImage{}
	for rows.Next() {
		err := rows.Scan(simsf.scanDest...)
		if err != nil {
			return nil, err
		}
		sim := skuimage.SKUImage{
			ID:    simsf.Receiver.SKUImage.ID,
			SKUID: simsf.Receiver.SKUImage.SKUID,
			Path:  simsf.Receiver.SKUImage.Path,
		}
		sims = append(sims, sim)
	}
	return sims, nil
}

func SKUImageSQLJoin(sf models.SQLFetcher, foreignKey string) *skuimage.SKUImage {
	sim := &skuimage.SKUImage{}
	dest := []interface{}{&sim.ID, &sim.Path, &sim.SKUID}
	sf.AddScanDest(dest)
	sf.AddJoins(" INNER JOIN sku_image ON sku_image.id = " + foreignKey)
	sf.AddFields(",sku_image.id,sku_image.path,sku_image.sku_id")
	return sim
}

func SQLSKUImageSQLJoin(sf models.SQLFetcher) *skuimage.SQLSKUImage {
	ssim := &skuimage.SQLSKUImage{}
	dest := []interface{}{&ssim.ID, &ssim.Path, &ssim.SKUID}
	sf.AddScanDest(dest)
	sf.AddJoins(" LEFT JOIN sku_image ON sku_image.sku_id = sku.id")
	sf.AddFields(",sku_image.id,sku_image.path,sku_image.sku_id")
	return ssim
}
