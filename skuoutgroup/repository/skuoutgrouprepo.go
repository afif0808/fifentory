package skuoutgrouprepo

import (
	"context"
	"fifentory/skuoutgroup"
)

type CreateSKUOutGroupFunc func(ctx context.Context, outgroup skuoutgroup.SKUOutGroup) (skuoutgroup.SKUOutGroup, error)

type GetSKUOutGroupsFunc func(ctx context.Context) ([]skuoutgroup.SKUOutGroup, error)
