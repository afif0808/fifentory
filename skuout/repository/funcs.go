package skuoutrepo

import (
	"context"
	"fifentory/skuout"
)

type CreateSKUOutFunc func(ctx context.Context, out skuout.SKUOut) (skuout.SKUOut, error)
type DeleteSKUOutByIDFunc func(ctx context.Context, id int64) error
type GetSKUOUtByIDFunc func(ctx context.Context, id int64) (*skuout.SKUOut, error)
type GetSKUOutsFunc func(ctx context.Context) ([]skuout.SKUOut, error)
type GetSKUOutsByGroupIDFunc func(ctx context.Context, groupID int64) ([]skuout.SKUOut, error)
type UpdateSKUOutFunc func(ctx context.Context, skuOut skuout.SKUOut) error
type DeleteSKUOutBySKUIDFunc func(ctx context.Context, skuID int64) error
type DeleteSKUOutByGroupIDFunc func(ctx context.Context, groupID int64) error
