package variantrepo

import (
	"context"
	"fifentory/options"
	variantvalue "fifentory/variant"
)

type GetVariantsFunc func(ctx context.Context, opts *options.Options) ([]variantvalue.Variant, error)
type CreateVariantFunc func(ctx context.Context, vv variantvalue.Variant) (variantvalue.Variant, error)
