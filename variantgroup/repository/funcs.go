package variantgrouprepo

import (
	"context"
	"fifentory/options"
	"fifentory/variantgroup"
)

type GetVariantGroupsFunc func(context.Context, *options.Options) ([]variantgroup.VariantGroup, error)
