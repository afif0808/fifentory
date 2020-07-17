package variantgroupresthandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	variantgrouprepo "fifentory/variantgroup/repository"
	variantgroupsqlrepo "fifentory/variantgroup/repository/sql"

	"net/http"
	"time"

	"github.com/labstack/echo"
)

func InjectVariantGroupRESTHandler(conn *sql.DB, ee *echo.Echo) {
	variantGroupSQLFetcher := variantgroupsqlrepo.NewVariantGroupSQLFetcher(conn)
	ee.GET("/variantgroups", GetVariantGroups(variantGroupSQLFetcher.Fetch))
}

func GetVariantGroups(getVariantGroups variantgrouprepo.GetVariantGroupsFunc) echo.HandlerFunc {
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
		vgs, err := getVariantGroups(ctx, nil)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		return ectx.JSON(http.StatusOK, vgs)
	}
}
