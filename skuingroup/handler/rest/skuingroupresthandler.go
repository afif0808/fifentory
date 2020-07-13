package skuingroupresthandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	skuinrepo "fifentory/skuin/repository"
	skuinsqlrepo "fifentory/skuin/repository/sql"
	skuingrouprepo "fifentory/skuingroup/repository"
	skuingroupsqlrepo "fifentory/skuingroup/repository/sql"
	stockrepo "fifentory/stock/repository"
	stocksqlrepo "fifentory/stock/repository/sql"
	suppliersqlrepo "fifentory/supplier/repository/sql"
	"log"

	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func InjectSKUInGroupRESTHandler(conn *sql.DB, ee *echo.Echo) {
	skuInGetter := skuinsqlrepo.NewSKUInGetter(conn)
	skuInGetter.WithSKU()
	skuInGetter.WithProduct()
	ee.GET("/skuingroups/:id/skuins", GetSKUInsByGroupID(skuInGetter.Get))

	skuInGroupFetcher := skuingroupsqlrepo.NewSKUInGroupFetcher(conn)

	sp := suppliersqlrepo.SupplierSQLJoin(&skuInGroupFetcher, "supplier_id")

	skuInGroupFetcher.SKUInGroup.Supplier = sp

	ee.GET("/skuingroups", GetSKUInGroups(skuInGroupFetcher.Fetch))

	deleteSKUIn := skuinsqlrepo.DeleteSKUIn(conn)
	deleteSKUInGroups := skuingroupsqlrepo.DeleteSKUInGroups(conn)
	substractStockQuantity := stocksqlrepo.SubstractStockQuantity(conn)
	ee.DELETE("/skuingroups/:id", DeleteSKUInGroups(
		deleteSKUInGroups,
		deleteSKUIn,
		substractStockQuantity,
		skuInGetter.Get))

}

func GetSKUInGroups(getSKUInGroups skuingrouprepo.GetSKUInGroupsFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		if _, isWithDeadline := ctx.Deadline(); isWithDeadline == false {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, time.Second*5)
			defer cancel()
		}
		skuInGroups, err := getSKUInGroups(ctx, nil)
		for _, v := range skuInGroups {
			log.Println(v)
		}

		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		return ectx.JSON(http.StatusOK, skuInGroups)
	}
}

func GetSKUInsByGroupID(getSKUIns skuinrepo.GetSKUInsFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		if _, isWithDeadline := ctx.Deadline(); isWithDeadline == false {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, time.Second*5)
			defer cancel()
		}
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		ft := options.Filter{
			By:       "sku_in_group_id",
			Operator: "=",
			Value:    id,
		}
		fts := []options.Filter{ft}
		opts := options.Options{Filters: fts}
		skuIns, err := getSKUIns(ctx, &opts)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, skuIns)
	}
}

func DeleteSKUInGroups(
	deleteSKUInGroups skuingrouprepo.DeleteSKUInGroupsFunc,
	deleteSKUIn skuinrepo.DeleteSKUInFunc,
	substractStockQuantity stockrepo.SubtractStockQuantityFunc,
	getSKUIns skuinrepo.GetSKUInsFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		if _, isWithDeadline := ctx.Deadline(); isWithDeadline == false {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, time.Second*5)
			defer cancel()
		}
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		skuInGroupsFilter := []options.Filter{options.Filter{
			By:       "sku_in_group.id",
			Operator: "=",
			Value:    id,
		}}
		err = deleteSKUInGroups(ctx, skuInGroupsFilter)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}

		skuInsFilters := []options.Filter{options.Filter{
			By:       "sku_in.sku_in_group_id",
			Operator: "=",
			Value:    id,
		}}
		skuIns, err := getSKUIns(ctx, &options.Options{Filters: skuInsFilters})
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		for _, v := range skuIns {
			stockFilters := []options.Filter{options.Filter{
				By:       "sku_stock.sku_id",
				Operator: "=",
				Value:    v.SKU.ID,
			}}
			err = substractStockQuantity(ctx, stockFilters, v.Quantity)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}

		}
		err = deleteSKUIn(ctx, skuInsFilters)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, "Succeeded deleting skuin group")
	}
}
