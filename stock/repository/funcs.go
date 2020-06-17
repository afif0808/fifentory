package stockrepo

import (
	"context"
	"fifentory/stock"
)

type CreateSKUStockFunc func(ctx context.Context, st stock.Stock) (stock.Stock, error)

type GetSKUStockBySKUID func(ctx context.Context, skuID int64) (*stock.Stock, error)

type UpdateSKUStockFunc func(ctx context.Context, st stock.Stock) (stock.Stock, error)

type SubtractStockQuantityBySKUIDFunc func(ctx context.Context, skuID int64, quantity int) error

type AddStockQuantityBySKUIDFunc func(ctx context.Context, skuID int64, quantity int) error
