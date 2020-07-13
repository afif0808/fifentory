package supplieresthandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/supplier"
	supplierrepo "fifentory/supplier/repository"
	suppliersqlrepo "fifentory/supplier/repository/sql"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func InjectSupplierRESTHandler(conn *sql.DB, ee *echo.Echo) {
	ssf := suppliersqlrepo.NewSupplierSQLFetcher(conn)
	ee.GET("/suppliers", GetSuppliers(ssf.Fetch))

	createSupplier := suppliersqlrepo.CreateSupplier(conn)
	ee.POST("/suppliers", CreateSupplier(createSupplier))
}

func GetSuppliers(getSuppliers supplierrepo.GetSuppliersFunc) echo.HandlerFunc {
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
		suppliers, err := getSuppliers(ctx, nil)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		return ectx.JSON(http.StatusOK, suppliers)
	}
}

func CreateSupplier(createSupplier supplierrepo.CreateSupplierFunc) echo.HandlerFunc {
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
		var post struct {
			Supplier *supplier.Supplier `json:"supplier"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{err.Error()})
		}
		*post.Supplier, err = createSupplier(ctx, *post.Supplier)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{err.Error()})
		}
		return ectx.JSON(http.StatusCreated, post.Supplier)

	}
}
