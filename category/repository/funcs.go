package categoryrepo

import (
	"context"
	"fifentory/category"
)

type GetProductCategoriesFunc func(ctx context.Context, productID int64) ([]category.Category, error)

type GetCategoryByIDFunc func(ctx context.Context, id int64) (category.Category, error)

type CreateProductCategoryJunctionFunc func(
	ctx context.Context,
	junction category.ProductCategoryJunction) (category.ProductCategoryJunction, error)
