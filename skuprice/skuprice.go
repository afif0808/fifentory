package skuprice

type SKUPrice struct {
	SKUID        int64 `json:"sku_id"`
	SellingPrice int   `json:"selling_price"`
	BuyingPrice  int   `json:"buying_price"`
}
