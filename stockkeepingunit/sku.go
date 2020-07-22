package stockkeepingunit

import (
	"fifentory/product"
	"fifentory/skuprice"
	"fifentory/stock"
	"time"
)

type StockKeepingUnit struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	ProductID int64     `json:"-"`
}

// sku with the foreign entities
type CompleteSKU struct {
	StockKeepingUnit
	Product product.Product `json:"product"`
	Stock   stock.Stock     `json:"stock"`
}

type SKU struct {
	ID        int64              `json:"id"`
	Code      string             `json:"code"`
	CreatedAt time.Time          `json:"created_at"`
	Product   *product.Product   `json:"product,omitempty" `
	Stock     *stock.Stock       `json:"stock,omitempty" `
	Price     *skuprice.SKUPrice `json:"price,omitempty"`
}
