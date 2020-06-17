package customerresthandler

import (
	"context"
	"database/sql"
	"errors"
	customerrepo "fifentory/customer/repository"
	customersqlrepo "fifentory/customer/repository/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InjectCustomerRESTHandler(conn *sql.DB, ee *echo.Echo) {
	getCustomers := customersqlrepo.GetCustomers(conn)
	ee.GET("/customers", GetCustomers(getCustomers))

	ee.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

}
func GetCustomers(getCustomers customerrepo.GetCustomersFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		customers, err := getCustomers(ctx)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, customers)
	}
}
func GetCustomerByID(getCustomer customerrepo.GetCustomerByIDFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		id, err := strconv.ParseInt((ectx.Param("id")), 64, 10)
		if err != nil {
			return ectx.JSON(http.StatusNotFound, errors.New("the user you requested was not found"))
		}
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		cus, err := getCustomer(ctx, id)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, cus)
	}
}
