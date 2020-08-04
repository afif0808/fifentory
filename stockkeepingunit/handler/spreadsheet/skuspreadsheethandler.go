package skuspreadsheethandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	productsqlrepo "fifentory/product/repository/sql"
	"fifentory/skuin"
	skuinrepo "fifentory/skuin/repository"
	skuinsqlrepo "fifentory/skuin/repository/sql"
	"fifentory/skuout"
	skuoutrepo "fifentory/skuout/repository"
	skuoutsqlrepo "fifentory/skuout/repository/sql"
	stocksqlrepo "fifentory/stock/repository/sql"
	"fifentory/stockkeepingunit"
	skurepo "fifentory/stockkeepingunit/repository"
	skusqlrepo "fifentory/stockkeepingunit/repository/sql"
	"net/http"
	"time"

	"github.com/tealeg/xlsx"

	"github.com/labstack/echo"
)

func InjectSKUSpreadSheetHandler(conn *sql.DB, ee *echo.Echo) {
	ssf := skusqlrepo.NewSKUSQLFetcher(conn)
	st := stocksqlrepo.SKUStockSQLJoin(&ssf)
	ssf.Receiver.Stock = st
	prod := productsqlrepo.ProductSQLJoin(&ssf, "sku.product_id")
	ssf.Receiver.Product = prod

	ssof := skuoutsqlrepo.NewSKUOutSQLFetcher(conn)

	ssif := skuinsqlrepo.NewSKUInGetter(conn)
	ee.GET("/accessories/report/", GetSKUs(ssf.Fetch, ssof.Fetch, ssif.Get))
}

func cancelSKUOut(sku stockkeepingunit.SKU, sos []skuout.SKUOut) stockkeepingunit.SKU {
	for _, so := range sos {
		sku.Stock.Quantity += so.Quantity
	}
	return sku
}
func cancelSKUIn(sku stockkeepingunit.SKU, sis []skuin.SKUIn) stockkeepingunit.SKU {
	for _, si := range sis {
		sku.Stock.Quantity -= si.Quantity
	}
	return sku
}

func GetSKUs(
	getSKUs skurepo.GetSKUsFunc,
	getSKUOuts skuoutrepo.GetSKUOutsFunc,
	getSKUIns skuinrepo.GetSKUInsFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		if _, isWithDeadline := ctx.Deadline(); !isWithDeadline {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, time.Second*5)
			defer cancel()
		}

		urlQuery := ectx.Request().URL.Query()

		f := xlsx.NewFile()
		sh, err := f.AddSheet("DATA")
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}

		skus, err := getSKUs(ctx, nil)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}

		if urlQuery["date"] != nil {
			date := urlQuery.Get("date")
			getSKUOutsByID := func(skuID int64) ([]skuout.SKUOut, error) {
				ft1 := options.Filter{
					By:       "sku_out.sku_id",
					Operator: "=",
					Value:    skuID,
				}
				ft2 := options.Filter{
					By:       "sku_out.date",
					Operator: ">",
					Value:    date,
				}
				return getSKUOuts(ctx, &options.Options{Filters: []options.Filter{ft1, ft2}})
			}
			getSKUInsByID := func(skuID int64) ([]skuin.SKUIn, error) {
				ft1 := options.Filter{
					By:       "sku_in.sku_id",
					Operator: "=",
					Value:    skuID,
				}
				ft2 := options.Filter{
					By:       "sku_in.date",
					Operator: ">",
					Value:    date,
				}
				return getSKUIns(ctx, &options.Options{Filters: []options.Filter{ft1, ft2}})
			}

			for i, s := range skus {
				sos, err := getSKUOutsByID(s.ID)
				if err != nil {
					if err != nil {
						return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
					}
				}
				skus[i] = cancelSKUOut(s, sos)

				sis, err := getSKUInsByID(s.ID)
				if err != nil {
					if err != nil {
						return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
					}
				}
				skus[i] = cancelSKUIn(s, sis)
			}
		}
		title := "DATA STOK AKSESORIS"
		fileName := "data_stok_aksesoris.xlsx"
		fileLocation := "/media/afif0808/data/go/src/fifentory/assets/spreadsheets/"

		if urlQuery["date"] != nil {
			title += " - " + urlQuery.Get("date")
			fileName = "data_stok_aksesoris" + "_" + urlQuery.Get("date") + ".xlsx"
		}

		firstRow := sh.AddRow()
		firstRowCell := firstRow.AddCell()
		firstRowCell.SetValue(title)
		firstRowCell.Merge(1, 0)

		headRow := sh.AddRow()
		headRow.AddCell().SetValue("NAMA")
		headRow.AddCell().SetValue("JUMLAH")

		for _, s := range skus {
			ro := sh.AddRow()
			ro.AddCell().SetValue(s.Product.Name)
			ro.AddCell().SetValue(s.Stock.Quantity)
		}

		err = f.Save(fileLocation + fileName)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		ectx.Response().Header().Set("Content-Disposition", "attachment; filename="+fileName)

		return ectx.File(fileLocation + fileName)
	}
}
