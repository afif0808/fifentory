package variant

import "fifentory/variantgroup"

type Variant struct {
	ID    int64                      `json:"id"`
	Value string                     `json:"value"`
	Group *variantgroup.VariantGroup `json:"group,omitempty"`
}
