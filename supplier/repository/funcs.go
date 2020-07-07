package supplierrepo

import (
	"context"
	"fifentory/options"
	"fifentory/supplier"
)

type GetSuppliersFunc func(ctx context.Context, opts *options.Options) ([]supplier.Supplier, error)
