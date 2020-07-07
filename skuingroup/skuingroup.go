package skuingroup

import (
	"fifentory/supplier"
	"time"
)

type SKUInGroup struct {
	ID       int64              `json:"id"`
	Date     time.Time          `json:"date"`
	Supplier *supplier.Supplier `json:"supplier,omitempty"`
}
