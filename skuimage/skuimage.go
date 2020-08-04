package skuimage

import (
	"database/sql"
)

type SKUImage struct {
	ID    int64  `json:"id"`
	Path  string `json:"path"`
	SKUID int64  `json:"sku_id"`
}

type SQLSKUImage struct {
	ID    sql.NullInt64
	Path  sql.NullString
	SKUID sql.NullInt64
}
