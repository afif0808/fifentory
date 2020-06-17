package skuoutgroup

import "time"

type SKUOutGroup struct {
	ID         int64     `json:"id"`
	Date       time.Time `json:"date"`
	CustomerID int64     `json:"customer_id"`
}
