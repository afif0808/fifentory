package skuingroup

import "time"

type SKUInGroup struct {
	ID         int64     `json:"id"`
	SupplierID int64     `json:"supplier_id"`
	Date       time.Time `json:"date"`
}
