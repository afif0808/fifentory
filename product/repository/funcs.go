package productrepo

import (
	"context"
	"fifentory/product"
)

type GetProductByIDFunc func(ctx context.Context, id int64) (*product.Product, error)
type CreateProductFunc func(ctx context.Context, prod product.Product) (product.Product, error)
type UpdateProductFunc func(ctx context.Context, prod product.Product) (product.Product, error)
