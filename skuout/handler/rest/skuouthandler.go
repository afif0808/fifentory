package skuoutresthandler

import (
	"context"
	"database/sql"
	"fifentory/skuout"
	skuoutrepo "fifentory/skuout/repository"
	skuoutsqlrepo "fifentory/skuout/repository/sql"
	"fifentory/skuoutgroup"
	skuoutgrouprepo "fifentory/skuoutgroup/repository"
	skuoutgroupsqlrepo "fifentory/skuoutgroup/repository/sql"
	stockrepo "fifentory/stock/repository"
	stocksqlrepo "fifentory/stock/repository/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InjectSKUOutHandler(conn *sql.DB, ee *echo.Echo) {
	createSKUOut := skuoutsqlrepo.CreateSKUOut(conn)
	subtractStock := stocksqlrepo.SubtractStockQuantityBySKUID(conn)
	createSKUOutGroup := skuoutgroupsqlrepo.CreateSKUOutGroup(conn)
	ee.POST("/skuouts", CreateSKUOut(subtractStock, createSKUOut, createSKUOutGroup))

	deleteSKUOutByID := skuoutsqlrepo.DeleteSKUOutByID(conn)
	getSKUOutByID := skuoutsqlrepo.GetSKUOUtByID(conn)
	addSKUStockQtyBySKUID := stocksqlrepo.AddStockQuantityBySKUID(conn)

	ee.DELETE("/skuouts/:id", DeleteSKUOutByID(deleteSKUOutByID, addSKUStockQtyBySKUID, getSKUOutByID))

	updateSKUOutByID := skuoutsqlrepo.UpdateSKUOutByID(conn)
	ee.POST("/skuouts/:id", UpdateSKUOutByID(
		updateSKUOutByID,
		getSKUOutByID,
		addSKUStockQtyBySKUID,
		subtractStock))

	ee.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

}

type ErrorResponse struct {
	Message string `json:"msg"`
}

type createSKUOutPost struct {
	SKUOuts     []skuout.SKUOut          `json:"sku_outs"`
	SKUOutGroup *skuoutgroup.SKUOutGroup `json:"sku_out_group"`
}

func CreateSKUOut(
	subractStock stockrepo.SubtractStockQuantityBySKUIDFunc,
	createSKUOut skuoutrepo.CreateSKUOutFunc,
	createSKUOutGroup skuoutgrouprepo.CreateSKUOutGroupFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		post := createSKUOutPost{}
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}

		if post.SKUOutGroup.ID == 0 {
			if post.SKUOutGroup.CustomerID == 0 {
				return ectx.JSON(http.StatusBadRequest, fmt.Errorf("Error : missing customer"))
			}
			if post.SKUOutGroup.Date == (time.Time{}) {
				post.SKUOutGroup.Date = time.Now().Local()
			}
			*post.SKUOutGroup, err = createSKUOutGroup(ctx, *post.SKUOutGroup)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
		}
		for i, out := range post.SKUOuts {
			out.GroupID = post.SKUOutGroup.ID
			if out.Date == (time.Time{}) {
				out.Date = time.Now().Local()
			}
			post.SKUOuts[i], err = createSKUOut(ctx, out)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			err = subractStock(ctx, out.SKUID, out.Quantity)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
		}
		return ectx.JSON(http.StatusOK, post)
	}
}

func DeleteSKUOutByID(
	deleteSKUOut skuoutrepo.DeleteSKUOutByIDFunc,
	addSKUStockQty stockrepo.AddStockQuantityBySKUIDFunc,
	getSKUOut skuoutrepo.GetSKUOUtByIDFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		}
		skuOut, err := getSKUOut(ctx, id)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}

		err = addSKUStockQty(ctx, skuOut.SKUID, skuOut.Quantity)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}

		err = deleteSKUOut(ctx, id)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}

		return ectx.JSON(http.StatusOK, id)
	}
}
func GetSKUOuts(
	getSKUOuts skuoutrepo.GetSKUOutsFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		return nil
	}
}

func UpdateSKUOutByID(
	updateSKUOut skuoutrepo.UpdateSKUOutFunc,
	getSKUOutByID skuoutrepo.GetSKUOUtByIDFunc,
	addStockQuantityBySKUID stockrepo.AddStockQuantityBySKUIDFunc,
	substractStockQuantityBySKUID stockrepo.SubtractStockQuantityBySKUIDFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		var post struct {
			SKUOut *skuout.SKUOut `json:"sku_out"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}

		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		post.SKUOut.ID = id

		skuOut, err := getSKUOutByID(ctx, id)
		diff := skuOut.Quantity - post.SKUOut.Quantity
		switch true {
		case diff > 0:
			err := addStockQuantityBySKUID(ctx, skuOut.SKUID, -diff)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}

			break
		case diff < 0:
			err := substractStockQuantityBySKUID(ctx, skuOut.SKUID, diff)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
		}

		err = updateSKUOut(ctx, *post.SKUOut)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}

		return ectx.JSON(http.StatusOK, post.SKUOut)
	}
}
