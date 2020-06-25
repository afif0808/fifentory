package skuingrouprepo

import (
	"context"
	"fifentory/skuingroup"
)

type CreateSKUInGroupFunc func(ctx context.Context, skuingroup skuingroup.SKUInGroup) (skuingroup.SKUInGroup, error)
