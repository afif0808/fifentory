package skuvariantrepo

import (
	"context"
	"fifentory/skuvariant"
)

type CreateSKUVariantFunc func(ctx context.Context, sv skuvariant.SKUVariant) (skuvariant.SKUVariant, error)
