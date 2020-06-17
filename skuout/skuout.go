package skuout

import "time"

type SKUOut struct {
	ID       int64     `json:"id"`
	Quantity int       `json:"qty"`
	SKUID    int64     `json:"sku_id"`
	Date     time.Time `json:"date"`
	GroupID  int64     `json:"group_id"`
}
