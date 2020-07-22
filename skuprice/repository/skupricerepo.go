package skupricerepo

import (
	"context"
	"fifentory/options"
	"fifentory/skuprice"
)

type GetSKUPricesFunc func(ctx context.Context, opts *options.Options) ([]skuprice.SKUPrice, error)
type CrateSKUPriceFunc func(ctx context.Context, sp skuprice.SKUPrice) (skuprice.SKUPrice, error)
type UpdateSKUPriceFunc func(ctx context.Context, sp skuprice.SKUPrice, fts []options.Filter) (skuprice.SKUPrice, error)
