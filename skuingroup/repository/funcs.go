package skuingrouprepo

import (
	"context"
	"fifentory/options"
	"fifentory/skuingroup"
)

type CreateSKUInGroupFunc func(ctx context.Context, skuingroup skuingroup.SKUInGroup) (skuingroup.SKUInGroup, error)
type GetSKUInGroupsFunc func(ctx context.Context, opts *options.Options) ([]skuingroup.SKUInGroup, error)
