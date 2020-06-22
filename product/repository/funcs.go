package productrepo

import (
	"context"
	"fifentory/options"
	"fifentory/product"
)

type GetProductByIDFunc func(ctx context.Context, id int64) (*product.Product, error)
type CreateProductFunc func(ctx context.Context, prod product.Product) (product.Product, error)
type UpdateProductFunc func(ctx context.Context, prod product.Product) (product.Product, error)
type GetProductsFunc func(ctx context.Context, opt *options.Options) ([]product.Product, error)
