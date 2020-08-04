package handler

import (
	"context"
	filerepo "fifentory/file/repository"
	"fifentory/models"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func UploadFile(uploadFile filerepo.UploadFileFunc, destpath string) echo.HandlerFunc {
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
		file, err := ectx.FormFile("file")
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{Message: err.Error()})
		}
		src, err := file.Open()
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		err = uploadFile(ctx, destpath+file.Filename, src)
		return ectx.JSON(http.StatusOK, nil)
	}
}
