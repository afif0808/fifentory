package skuin

import (
	"fifentory/skuingroup"
	"fifentory/stockkeepingunit"
	"time"
)

type SKUIn struct {
	ID       int64                  `json:"id"`
	Quantity int                    `json:"qty"`
	Date     time.Time              `json:"date"`
	SKU      *stockkeepingunit.SKU  `json:"sku,omitempty"`
	Group    *skuingroup.SKUInGroup `json:"group,omitempty"`
}
