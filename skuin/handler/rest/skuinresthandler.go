package skuinresthandler

import (
	"context"
	"database/sql"
	"errors"
	"fifentory/models"
	"fifentory/options"
	"fifentory/skuin"
	skuinrepo "fifentory/skuin/repository"
	skuinsqlrepo "fifentory/skuin/repository/sql"
	"fifentory/skuingroup"
	skuingrouprepo "fifentory/skuingroup/repository"
	skuingroupsqlrepo "fifentory/skuingroup/repository/sql"
	stockrepo "fifentory/stock/repository"
	stocksqlrepo "fifentory/stock/repository/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func InjectSKUInRESTHandler(conn *sql.DB, ee *echo.Echo) {
	createSKUIn := skuinsqlrepo.CreateSKUIn(conn)
	createSKUInGroup := skuingroupsqlrepo.CreateSKUInGroup(conn)
	addStockQuantityBySKUID := stocksqlrepo.AddStockQuantityBySKUID(conn)
	ee.POST("/skuins", CreateSKUIn(createSKUInGroup, createSKUIn, addStockQuantityBySKUID))

	skuInGetter := skuinsqlrepo.NewSKUInGetter(conn)
	skuInGetter.WithSKU()
	skuInGetter.WithProduct()
	skuInGetter.WithGroup()
	ee.GET("/skuins", GetSKUIns(skuInGetter.Get))

	substractStockQuantity := stocksqlrepo.SubstractStockQuantity(conn)
	addStockQuantity := stocksqlrepo.AddStockQuantity(conn)
	deleteSKUIn := skuinsqlrepo.DeleteSKUIn(conn)
	updateSKUIn := skuinsqlrepo.UpdateSKUIn(conn)
	ee.DELETE("/skuins/:id", DeleteSKUIn(deleteSKUIn, substractStockQuantity, skuInGetter.Get))
	ee.POST("/skuins/:id", UpdateSKUIn(substractStockQuantity, addStockQuantity, skuInGetter.Get, updateSKUIn))

}

func CreateSKUIn(
	createSKUInGroup skuingrouprepo.CreateSKUInGroupFunc,
	createSKUIn skuinrepo.CreateSKUInFunc,
	addStockQuantityBySKUID stockrepo.AddStockQuantityBySKUIDFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
		}
		var post struct {
			SKUIns []skuin.SKUIn          `json:"sku_ins"`
			Group  *skuingroup.SKUInGroup `json:"sku_in_group"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{Message: err.Error()})
		}
		if post.Group.ID == 0 {
			*post.Group, err = createSKUInGroup(ctx, *post.Group)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			if post.Group.Date == (time.Time{}) {
				post.Group.Date = time.Now()
			}
		}
		for i := range post.SKUIns {
			if post.SKUIns[i].Date == (time.Time{}) {
				post.SKUIns[i].Date = time.Now()
			}

			if post.SKUIns[i].Group == nil {
				post.SKUIns[i].Group = &skuingroup.SKUInGroup{}
			}

			post.SKUIns[i].Group.ID = post.Group.ID
			post.SKUIns[i], err = createSKUIn(ctx, post.SKUIns[i])

			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			err := addStockQuantityBySKUID(ctx, post.SKUIns[i].SKU.ID, post.SKUIns[i].Quantity)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
		}
		return ectx.JSON(http.StatusCreated, post)
	}
}

func GetSKUIns(getSKUIns skuinrepo.GetSKUInsFunc) echo.HandlerFunc {
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

		skuIns, err := getSKUIns(ctx, nil)
		if err != nil {
			log.Println(err)
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, skuIns)
	}
}

func DeleteSKUIn(
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
			return ectx.JSON(http.StatusNotFound, models.RESTErrorResponse{Message: "skuin not found"})
		}
		skuInFilters := []options.Filter{options.Filter{
			By:       "sku_in.id",
			Value:    id,
			Operator: "=",
		}}

		skuIns, err := getSKUIns(ctx, &options.Options{Filters: skuInFilters})
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		if len(skuIns) <= 0 || skuIns[0].SKU == nil {
			return ectx.JSON(http.StatusInternalServerError, errors.New("cannot delete skuin , it could be the skuin not found or internal error"))

		}

		err = deleteSKUIn(ctx, skuInFilters)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		substractFilter := []options.Filter{options.Filter{
			By:       "sku_id",
			Value:    skuIns[0].SKU.ID,
			Operator: "=",
		}}
		err = substractStockQuantity(ctx, substractFilter, skuIns[0].Quantity)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		return ectx.JSON(http.StatusOK, "Succeded deleting the skuin")
	}
}

func UpdateSKUIn(
	substractStockQuantity stockrepo.SubtractStockQuantityFunc,
	addStockQuantity stockrepo.AddStockQuantityFunc,
	getSKUIns skuinrepo.GetSKUInsFunc,
	updateSKUIn skuinrepo.UpdateSKUInFunc,
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
			SKUIn skuin.SKUIn `json:"sku_in"`
		}

		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusNotFound, models.RESTErrorResponse{Message: "skuin not found"})
		}

		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusNotFound, models.RESTErrorResponse{Message: "skuin not found"})
		}

		skuInFilters := []options.Filter{options.Filter{
			By:       "sku_in.id",
			Value:    id,
			Operator: "=",
		}}

		skuIns, err := getSKUIns(ctx, &options.Options{Filters: skuInFilters})
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		if len(skuIns) <= 0 || skuIns[0].SKU == nil {
			return ectx.JSON(http.StatusInternalServerError, errors.New("cannot update skuin , it could be the skuin not found or internal error"))
		}

		post.SKUIn, err = updateSKUIn(ctx, post.SKUIn, skuInFilters)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}

		stockChangeFilter := []options.Filter{options.Filter{
			By:       "sku_id",
			Operator: "=",
			Value:    skuIns[0].SKU.ID,
		}}
		if post.SKUIn.Quantity > skuIns[0].Quantity {
			err = addStockQuantity(ctx, stockChangeFilter, post.SKUIn.Quantity-skuIns[0].Quantity)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
			}

		} else if post.SKUIn.Quantity < skuIns[0].Quantity {
			err = substractStockQuantity(ctx, stockChangeFilter, skuIns[0].Quantity-post.SKUIn.Quantity)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
			}

		}

		return ectx.JSON(http.StatusOK, post)
	}
}
