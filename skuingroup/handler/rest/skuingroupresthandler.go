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
	"fifentory/supplier"
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
	sp := supplier.Supplier{}
	suppliersqlrepo.SupplierSQLJoin(&skuInGroupFetcher, &sp, "supplier_id")
	skuInGroupFetcher.SKUInGroup.Supplier = &sp
	ee.GET("/skuingroups", GetSKUInGroups(skuInGroupFetcher.Fetch))

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
