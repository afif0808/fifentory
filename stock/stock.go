package stock

type Stock struct {
	ID              int64 `json:"id"`
	Quantity        int   `json:"qty"`
	MinimumQuantity int   `json:"minimum_qty"`
	SKUID           int64 `json:"-"`
}
