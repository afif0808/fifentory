package supplieresthandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	supplierrepo "fifentory/supplier/repository"
	suppliersqlrepo "fifentory/supplier/repository/sql"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func InjectSupplierRESTHandler(conn *sql.DB, ee *echo.Echo) {
	ssf := suppliersqlrepo.NewSupplierSQLFetcher(conn)
	ee.GET("/suppliers", GetSuppliers(ssf.Fetch))
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
