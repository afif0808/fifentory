package skuimagerepo

import (
	"context"
	"fifentory/options"
	"fifentory/skuimage"
)

type GetSKUImagesFunc func(ctx context.Context, opt *options.Options) ([]skuimage.SKUImage, error)
type CreateSKUImageFunc func(ctx context.Context, sim skuimage.SKUImage) (skuimage.SKUImage, error)
type DeleteSKUImagesFunc func(ctx context.Context , fts []options.Filter) error 