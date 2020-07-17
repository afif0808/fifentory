package variantresthandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/variant"
	variantrepo "fifentory/variant/repository"
	variantsqlrepo "fifentory/variant/repository/sql"

	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func InjectVariantRESTHandler(conn *sql.DB, ee *echo.Echo) {
	vvsf := variantsqlrepo.NewVariantSQLFetcher(conn)
	ee.GET("/variants", GetVariants(vvsf.Fetch))
	createVariant := variantsqlrepo.CreateVariant(conn)
	ee.POST("/variants", CreateVariants(createVariant))
}

func GetVariants(getVariants variantrepo.GetVariantsFunc) echo.HandlerFunc {
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

		urlQuery := ectx.Request().URL.Query()

		fts := []options.Filter{}
		if urlQuery.Get("group") != "" {
			variantOptionID, err := strconv.ParseInt(urlQuery.Get("group"), 10, 64)
			if err != nil {
				return ectx.JSON(
					http.StatusBadRequest,
					models.RESTErrorResponse{
						Message: "Error : no variant with given value was found ",
					},
				)
			}
			fts = append(fts, options.Filter{
				By:       "variant_group_id",
				Operator: "=",
				Value:    variantOptionID,
			})
		}

		opts := &options.Options{Filters: fts}
		variantValues, err := getVariants(ctx, opts)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		return ectx.JSON(http.StatusOK, variantValues)
	}
}

func CreateVariants(
	createVariant variantrepo.CreateVariantFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		if _, isWithDeadline := ctx.Deadline(); !isWithDeadline {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, time.Second*1)
			defer cancel()
		}
		var post struct {
			Variants []variant.Variant `json:"variants"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, models.RESTErrorResponse{Message: err.Error()})
		}
		for _, vv := range post.Variants {
			vv, err = createVariant(ctx, vv)
		}
		return ectx.JSON(http.StatusCreated, post.Variants)
	}
}
