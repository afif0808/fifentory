package variantgroupsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/variantgroup"
	"log"
)

type receiver struct {
	VariantGroup *variantgroup.VariantGroup
}

type VariantGroupSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *receiver
	conn     *sql.DB
}

func NewVariantGroupSQLFetcher(conn *sql.DB) VariantGroupSQLFetcher {
	vasf := VariantGroupSQLFetcher{
		fields: "variant_group.id,variant_group.name",
		Receiver: &receiver{
			VariantGroup: &variantgroup.VariantGroup{},
		},
		conn: conn,
	}
	vasf.scanDest = []interface{}{&vasf.Receiver.VariantGroup.ID, &vasf.Receiver.VariantGroup.Name}
	return vasf
}

func (vgsf *VariantGroupSQLFetcher) Fetch(ctx context.Context, opts *options.Options) ([]variantgroup.VariantGroup, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	rows, err := vgsf.conn.QueryContext(ctx, "SELECT "+vgsf.fields+" FROM "+variantGroupTable+" "+vgsf.joins+" "+optionsQuery, optionsArgs...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	vgs := []variantgroup.VariantGroup{}
	for rows.Next() {
		err = rows.Scan(vgsf.scanDest...)
		if err != nil {
			return nil, err
		}
		vg := variantgroup.VariantGroup{
			ID:   vgsf.Receiver.VariantGroup.ID,
			Name: vgsf.Receiver.VariantGroup.Name,
		}
		vgs = append(vgs, vg)
	}
	return vgs, nil
}
