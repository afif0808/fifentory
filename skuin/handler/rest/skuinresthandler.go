package skuinresthandler

import (
	"context"
	"database/sql"
	"fifentory/skuin"
	skuinrepo "fifentory/skuin/repository"
	skuinsqlrepo "fifentory/skuin/repository/sql"
	"fifentory/skuingroup"
	skuingrouprepo "fifentory/skuingroup/repository"
	skuingroupsqlrepo "fifentory/skuingroup/repository/sql"
	stockrepo "fifentory/stock/repository"
	stocksqlrepo "fifentory/stock/repository/sql"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func InjectSKUInRESTHandler(conn *sql.DB, ee *echo.Echo) {
	createSKUIn := skuinsqlrepo.CreateSKUIn(conn)
	createSKUInGroup := skuingroupsqlrepo.CreateSKUInGroup(conn)
	subtractStockQuantityBySKUID := stocksqlrepo.SubtractStockQuantityBySKUID(conn)
	ee.POST("/skuins", CreateSKUIn(createSKUInGroup, createSKUIn, subtractStockQuantityBySKUID))
}

func CreateSKUIn(
	createSKUInGroup skuingrouprepo.CreateSKUInGroupFunc,
	createSKUIn skuinrepo.CreateSKUInFunc,
	subtractStockQuantityBySKUID stockrepo.SubtractStockQuantityBySKUIDFunc,
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
			return ectx.JSON(http.StatusBadRequest, err)
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
			post.SKUIns[i].GroupID = post.Group.ID
			post.SKUIns[i], err = createSKUIn(ctx, post.SKUIns[i])

			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			err := subtractStockQuantityBySKUID(ctx, post.SKUIns[i].SKUID, post.SKUIns[i].Quantity)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
		}
		return ectx.JSON(http.StatusCreated, post)
	}
}
