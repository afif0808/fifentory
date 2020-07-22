package skupriceresthandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/skuprice"
	skupricerepo "fifentory/skuprice/repository"
	skupricesqlrepo "fifentory/skuprice/repository/sql"
	"log"

	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func InjectSKUPriceRESTHandler(conn *sql.DB, ee *echo.Echo) {
	updateSKUPrice := skupricesqlrepo.UpdateSKUPrice(conn)
	ee.POST("/skuprices/:skuID", UpdateSKUPrice(updateSKUPrice))
}

func UpdateSKUPrice(updateSKUPrice skupricerepo.UpdateSKUPriceFunc) echo.HandlerFunc {
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
			SKUPrice *skuprice.SKUPrice `json:"sku_price"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(
				http.StatusBadRequest,
				models.RESTErrorResponse{
					Message: err.Error(),
				})

		}
		skuID, err := strconv.ParseInt(ectx.Param("skuID"), 10, 64)
		if err != nil {
			log.Println(err)
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{Message: "Error : what are you looking for is not found"})
		}

		post.SKUPrice.SKUID = skuID

		*post.SKUPrice, err = updateSKUPrice(ctx, *post.SKUPrice, []options.Filter{options.Filter{
			By:       "sku_id",
			Operator: "=",
			Value:    skuID,
		}})
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		return ectx.JSON(http.StatusOK, post.SKUPrice)
	}
}
