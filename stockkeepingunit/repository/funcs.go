package skurepo

import (
	"context"
	"fifentory/options"
	"fifentory/stockkeepingunit"
)

type CreateSKUFunc func(ctx context.Context, sku stockkeepingunit.StockKeepingUnit) (stockkeepingunit.StockKeepingUnit, error)
type GetSKUs func(ctx context.Context, opts *options.Options) ([]stockkeepingunit.StockKeepingUnit, error)
type GetSKUsFunc func(ctx context.Context, opts *options.Options) ([]stockkeepingunit.SKU , error)
type GetSKUByIDFunc func(ctx context.Context, id int64) (*stockkeepingunit.StockKeepingUnit, error)
type DeleteSKUByID func(ctx context.Context, id int64) error
type UpdateSKUFunc func(ctx context.Context, sku stockkeepingunit.StockKeepingUnit) (stockkeepingunit.StockKeepingUnit, error)
