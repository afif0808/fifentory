package variantsqlrepo

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/variant"
	"fifentory/variantgroup"
	"log"
)

type receiver struct {
	Variant      *variant.Variant
	VariantGroup *variantgroup.VariantGroup
}

type VariantSQLFetcher struct {
	joins    string
	fields   string
	scanDest []interface{}
	Receiver *receiver
	conn     *sql.DB
}

func NewVariantSQLFetcher(conn *sql.DB) VariantSQLFetcher {
	vasf := VariantSQLFetcher{
		fields: "variant.id,variant.value",
		Receiver: &receiver{
			Variant: &variant.Variant{},
		},
		conn: conn,
	}
	vasf.scanDest = []interface{}{&vasf.Receiver.Variant.ID, &vasf.Receiver.Variant.Value}
	return vasf
}

func (vasf *VariantSQLFetcher) AddJoins(joins string) {
	vasf.joins += joins
}
func (vasf *VariantSQLFetcher) AddFields(fields string) {
	vasf.fields += fields
}
func (vasf *VariantSQLFetcher) AddScanDest(dest []interface{}) {
	vasf.scanDest = append(vasf.scanDest, dest...)
}

func (vasf *VariantSQLFetcher) Fetch(ctx context.Context, opts *options.Options) ([]variant.Variant, error) {
	optionsQuery, optionsArgs := options.ParseOptionsToSQLQuery(opts)
	query := "SELECT " + vasf.fields + " FROM " + variantValueTable + " " + vasf.joins + " " + optionsQuery

	rows, err := vasf.conn.QueryContext(ctx, query, optionsArgs...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	vas := []variant.Variant{}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(vasf.scanDest...)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		va := variant.Variant{
			ID:    vasf.Receiver.Variant.ID,
			Value: vasf.Receiver.Variant.Value,
		}
		if vasf.Receiver.VariantGroup != nil {
			vo := variantgroup.VariantGroup{
				ID:   vasf.Receiver.VariantGroup.ID,
				Name: vasf.Receiver.VariantGroup.Name,
			}
			va.Group = &vo
		}
		vas = append(vas, va)
	}
	return vas, nil
}

func VariantSQLJoin(sf models.SQLFetcher, foreignKey string) *variant.Variant {

	return nil
}
