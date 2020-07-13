package customerresthandler

import (
	"context"
	"database/sql"
	"errors"
	"fifentory/customer"
	customerrepo "fifentory/customer/repository"
	customersqlrepo "fifentory/customer/repository/sql"
	"fifentory/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InjectCustomerRESTHandler(conn *sql.DB, ee *echo.Echo) {
	getCustomers := customersqlrepo.GetCustomers(conn)
	ee.GET("/customers", GetCustomers(getCustomers))
	createCustomer := customersqlrepo.CreateCustomer(conn)
	ee.POST("/customers", CreateCustomer(createCustomer))
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

func CreateCustomer(
	createCustomer customerrepo.CreateCustomerFunc,
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
		var post struct {
			Customer *customer.Customer `json:"customer"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{Message: err.Error()})
		}
		*post.Customer, err = createCustomer(ctx, *post.Customer)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		return ectx.JSON(http.StatusCreated, post.Customer)
	}
}
