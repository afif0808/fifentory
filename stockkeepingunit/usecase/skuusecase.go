package skuusecase

import (
	"context"
	"fifentory/product"
	productrepo "fifentory/product/repository"
	stockrepo "fifentory/stock/repository"
	"fifentory/stockkeepingunit"
	skurepo "fifentory/stockkeepingunit/repository"
)

func GetCompleteSKUs(
	getSKUs skurepo.GetSKUs,
	getStock stockrepo.GetSKUStockBySKUID,
	getProduct productrepo.GetProductByIDFunc,
) GetCompleteSKUsFunc {
	return func(ctx context.Context) ([]stockkeepingunit.CompleteSKU, error) {
		skus, err := getSKUs(ctx)
		if err != nil {
			return nil, err
		}
		skusProduct := map[int64]*product.Product{}
		completeSKUs := []stockkeepingunit.CompleteSKU{}
		for _, sku := range skus {

			skusProduct[sku.ProductID] = nil

			csku := stockkeepingunit.CompleteSKU{}
			st, err := getStock(ctx, sku.ID)
			if err != nil {
				return nil, err
			}
			csku.StockKeepingUnit = sku
			csku.Stock = *st
			completeSKUs = append(completeSKUs, csku)
		}
		for i := range skusProduct {
			skusProduct[i], err = getProduct(ctx, i)
			if err != nil {
				return nil, err
			}
		}
		for i, csku := range completeSKUs {
			completeSKUs[i].Product = *skusProduct[csku.StockKeepingUnit.ProductID]
		}
		return completeSKUs, nil
	}
}

func GetCompleteSKUByID(
	getSKU skurepo.GetSKUByIDFunc,
	getStock stockrepo.GetSKUStockBySKUID,
	getProduct productrepo.GetProductByIDFunc,
) GetCompleteSKUByIDFunc {
	return func(ctx context.Context, id int64) (*stockkeepingunit.CompleteSKU, error) {
		sku, err := getSKU(ctx, id)
		if err != nil {
			return nil, err
		}

		stock, err := getStock(ctx, sku.ID)
		if err != nil {
			return nil, err
		}

		prod, err := getProduct(ctx, sku.ProductID)

		if err != nil {
			return nil, err
		}

		completeSKU := stockkeepingunit.CompleteSKU{StockKeepingUnit: *sku, Stock: *stock, Product: *prod}
		return &completeSKU, nil
	}
}
