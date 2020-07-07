package skuinrepo

import (
	"context"
	"fifentory/options"
	"fifentory/skuin"
)

type CreateSKUInFunc func(ctx context.Context, in skuin.SKUIn) (skuin.SKUIn, error)
type GetSKUInsFunc func(ctx context.Context, opts *options.Options) ([]skuin.SKUIn, error)
type DeleteSKUInFunc func(ctx context.Context, fts []options.Filter) error
type UpdateSKUInFunc func(ctx context.Context, in skuin.SKUIn, fts []options.Filter) (skuin.SKUIn, error)
