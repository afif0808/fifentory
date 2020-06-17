package stockresthandler

import (
	"context"
	"database/sql"
	"fifentory/stock"
	stockrepo "fifentory/stock/repository"
	stocksqlrepo "fifentory/stock/repository/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func InjectStockRESTHandler(conn *sql.DB, ee *echo.Echo) {
	updateSKUStockById := stocksqlrepo.UpdateSKUStockByID(conn)
	ee.POST("/stocks/:id", UpdateSKUStockById(updateSKUStockById))
}

func UpdateSKUStockById(updateSKUStock stockrepo.UpdateSKUStockFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {

		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		var post struct {
			Stock stock.Stock `json:"stock"`
		}

		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		post.Stock.ID = id
		st, err := updateSKUStock(ctx, post.Stock)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, st)
	}
}
