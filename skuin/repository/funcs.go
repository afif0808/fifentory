package skuinrepo

import (
	"context"
	"fifentory/skuin"
)

type CreateSKUInFunc func(ctx context.Context, in skuin.SKUIn) (skuin.SKUIn, error)
