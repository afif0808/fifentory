package imagehandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	skuimagerepo "fifentory/skuimage/repository"
	skuimagesqlrepo "fifentory/skuimage/repository/sql"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func InjectImagesHandler(conn *sql.DB, ee *echo.Echo) {
	ee.POST("/skuimages/upload", UploadImage())
	simsf := skuimagesqlrepo.NewSKUImageSQLFetcher(conn)
	ee.GET("/skuimages/image/:id", GetSKUImage(simsf.Fetch))
}

func GetSKUImage(getSKUImages skuimagerepo.GetSKUImagesFunc) echo.HandlerFunc {
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
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{Message: err.Error()})
		}
		opts := options.Options{}
		opts.Filters = []options.Filter{options.Filter{
			By:       "id",
			Operator: "=",
			Value:    id,
		}}
		sims, err := getSKUImages(ctx, &opts)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		if len(sims) < 1 {
			return ectx.JSON(http.StatusNotFound, models.RESTErrorResponse{Message: "sku image with given  not found"})
		}
		// img, err := os.Open(sims[0].Path)
		// if err != nil {
		// 	return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		// }
		return ectx.File(sims[0].Path)
	}
}

func UploadImage() echo.HandlerFunc {
	return func(ectx echo.Context) error {
		file, err := ectx.FormFile("image")
		if err != nil {
			log.Println("nol", err)
			return ectx.JSON(http.StatusBadRequest, err.Error())
		}
		src, err := file.Open()
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err.Error())
		}
		defer src.Close()

		path := "/media/afif0808/data/go/src/fifentory/assets/images/" + strconv.Itoa(int(time.Now().UnixNano()))
		dst, err := os.Create(path)
		if err != nil {
			log.Println("satu", err)
			return ectx.JSON(http.StatusInternalServerError, err.Error())
		}

		_, err = io.Copy(dst, src)

		if err != nil {
			log.Println("dua", err)
			return ectx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ectx.JSON(
			http.StatusOK,
			struct {
				Path string `json:"path"`
			}{Path: path},
		)
	}
}
