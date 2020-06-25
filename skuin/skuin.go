package skuin

import "time"

type SKUIn struct {
	ID       int64     `json:"id"`
	Quantity int       `json:"qty"`
	Date     time.Time `json:"date"`
	SKUID    int64     `json:"sku_id"`
	GroupID  int64     `json:"group_id"`
}
