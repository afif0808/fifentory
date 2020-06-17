package stock

type Stock struct {
	ID       int64 `json:"id"`
	Quantity int   `json:"qty"`
	SKUID    int64 `json:"-"`
}
