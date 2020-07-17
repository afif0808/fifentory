package skuvariant

type SKUVariant struct {
	ID             int64 `json:"id"`
	SKUID          int64 `json:"sku_id"`
	VariantValueID int64 `json:"variant_value_id"`
}
