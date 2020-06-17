package skuusecase

import (
	"context"
	"fifentory/stockkeepingunit"
)

type GetCompleteSKUsFunc func(context.Context) ([]stockkeepingunit.CompleteSKU, error)

type GetCompleteSKUByIDFunc func(ctx context.Context, id int64) (*stockkeepingunit.CompleteSKU, error)
