package skuimageresthandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/skuimage"
	skuimagesqlrepo "fifentory/skuimage/repository/sql"
	"strconv"

	skuimagerepo "fifentory/skuimage/repository"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func InjectSKUImageRESTHandler(conn *sql.DB, ee *echo.Echo) {
	createSKUImage := skuimagesqlrepo.CreateSKUImage(conn)
	ee.POST("/skuimages", CreateSKUImages(createSKUImage))

	simsf := skuimagesqlrepo.NewSKUImageSQLFetcher(conn)
	ee.GET("/skuimages/:id", GetSKUImageImage(simsf.Fetch))
	deleteSKUImages := skuimagesqlrepo.DeleteSKUImages(conn)
	ee.DELETE("/skuimages/:id", DeleteSKUImageByID(deleteSKUImages))

}

func GetSKUImageImage(getSKUImages skuimagerepo.GetSKUImagesFunc) echo.HandlerFunc {
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
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{Message: "no sku image found with given id"})
		}
		sims, err := getSKUImages(ctx, &options.Options{Filters: []options.Filter{options.Filter{
			By:       "id",
			Operator: "=",
			Value:    id,
		}}})
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		if len(sims) <= 0 {
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{Message: "no sku image found with given id"})
		}
		ectx.File(sims[0].Path)
		return nil
	}
}

func CreateSKUImages(createSKUImage skuimagerepo.CreateSKUImageFunc) echo.HandlerFunc {
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
			SKUImages []skuimage.SKUImage `json:"sku_images"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusOK, models.RESTErrorResponse{Message: err.Error()})
		}
		for i, sim := range post.SKUImages {
			sim, err := createSKUImage(ctx, sim)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
			}
			post.SKUImages[i] = sim
		}
		return ectx.JSON(http.StatusCreated, post.SKUImages)
	}
}

func DeleteSKUImageByID(deleteSKUImages skuimagerepo.DeleteSKUImagesFunc) echo.HandlerFunc {
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
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: "No sku image was found with given id"})
		}
		fts := []options.Filter{
			options.Filter{
				By:       "id",
				Operator: "=",
				Value:    id,
			},
		}
		err = deleteSKUImages(ctx, fts)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: "Failed to delete"})
		}
		return ectx.JSON(http.StatusOK, "Deleting sucess")
	}
}
