package stock

type Stock struct {
	ID         int64 `json:"id"`
	Quantity   int   `json:"qty"`
	MinimumQty int   `json:"minimun_qty"`
	SKUID      int64 `json:"-"`
}
