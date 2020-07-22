package stock

type Stock struct {
	Quantity        int   `json:"qty"`
	MinimumQuantity int   `json:"minimum_qty"`
	SKUID           int64 `json:"sku_id"`
}
