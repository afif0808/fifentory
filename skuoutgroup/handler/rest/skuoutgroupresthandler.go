package skuoutgroupresthandler

import (
	"context"
	"database/sql"
	"fifentory/customer"
	customerrepo "fifentory/customer/repository"
	customersqlrepo "fifentory/customer/repository/sql"
	"fifentory/product"
	productrepo "fifentory/product/repository"
	productsqlrepo "fifentory/product/repository/sql"
	"fifentory/skuout"
	skuoutrepo "fifentory/skuout/repository"
	skuoutsqlrepo "fifentory/skuout/repository/sql"
	"fifentory/skuoutgroup"
	skuoutgrouprepo "fifentory/skuoutgroup/repository"
	skuoutgroupsqlrepo "fifentory/skuoutgroup/repository/sql"
	"fifentory/stockkeepingunit"
	skurepo "fifentory/stockkeepingunit/repository"
	skusqlrepo "fifentory/stockkeepingunit/repository/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func InjectSKUOutGroupRESTHandler(conn *sql.DB, ee *echo.Echo) {
	getCustomer := customersqlrepo.GetCustomer(conn)
	getSKUOutGroups := skuoutgroupsqlrepo.GetSKUOutGroups(conn)
	ee.GET("/skuoutgroups", GetSKUOutGroups(getCustomer, getSKUOutGroups))

	getSKUOutsByGroupID := skuoutsqlrepo.GetSKUOutsByGroupId(conn)
	getSKUByID := skusqlrepo.GetSKUByID(conn)
	getProductByID := productsqlrepo.GetProductByID(conn)

	ee.GET("/skuoutgroups/:id/skuouts",
		GetSKUOutsByGroupID(getSKUOutsByGroupID, getProductByID, getSKUByID))

}

func GetSKUOutGroups(
	getCustomer customerrepo.GetCustomerByIDFunc,
	getSKUOutGroups skuoutgrouprepo.GetSKUOutGroupsFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		outgroups, err := getSKUOutGroups(ctx)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		customers := map[int64]*customer.Customer{}
		for _, og := range outgroups {
			customers[og.CustomerID] = nil
		}
		for cusId := range customers {
			customers[cusId], err = getCustomer(ctx, cusId)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}

		}
		log.Println(customers, err)
		type Response struct {
			skuoutgroup.SKUOutGroup
			Customer *customer.Customer `json:"customer"`
		}
		resp := []Response{}
		for _, og := range outgroups {
			r := Response{SKUOutGroup: og, Customer: customers[og.CustomerID]}
			resp = append(resp, r)
		}

		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, resp)

	}
}

type completeSKUOut struct {
	skuout.SKUOut
	SKU     stockkeepingunit.StockKeepingUnit `json:"sku"`
	Product product.Product                   `json:"product"`
}

func GetSKUOutsByGroupID(
	getSKUOuts skuoutrepo.GetSKUOutsByGroupIDFunc,
	getProduct productrepo.GetProductByIDFunc,
	getSKU skurepo.GetSKUByIDFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		groupID, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusNotFound, err)
		}
		skuOuts, err := getSKUOuts(ctx, groupID)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		completeSKUOuts := []completeSKUOut{}
		for _, skuOut := range skuOuts {

			sku, err := getSKU(ctx, skuOut.SKUID)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			log.Println(*sku)
			prod, err := getProduct(ctx, sku.ProductID)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			completeSKUOut := completeSKUOut{SKUOut: skuOut, SKU: *sku, Product: *prod}

			completeSKUOuts = append(completeSKUOuts, completeSKUOut)
		}
		return ectx.JSON(http.StatusOK, completeSKUOuts)
	}
}
