package skuresthandler

import (
	"context"
	"database/sql"
	"fifentory/models"
	"fifentory/options"
	"fifentory/product"
	productrepo "fifentory/product/repository"
	productsqlrepo "fifentory/product/repository/sql"
	skuoutrepo "fifentory/skuout/repository"
	skuoutsqlrepo "fifentory/skuout/repository/sql"
	"fifentory/skuprice"
	skupricerepo "fifentory/skuprice/repository"
	skupricesqlrepo "fifentory/skuprice/repository/sql"
	"fifentory/skuvariant"
	skuvariantrepo "fifentory/skuvariant/repository"
	skuvariantsqlrepo "fifentory/skuvariant/repository/sql"
	"fifentory/stock"
	stockrepo "fifentory/stock/repository"
	stocksqlrepo "fifentory/stock/repository/sql"
	"fifentory/stockkeepingunit"
	skusqlrepo "fifentory/stockkeepingunit/repository/sql"
	"log"
	"strconv"

	"gopkg.in/go-playground/validator.v9"

	skurepo "fifentory/stockkeepingunit/repository"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InjectSKURESTHandler(conn *sql.DB, ee *echo.Echo) {
	// createProductCategoryJunction := categorysqlrepo.CreateProductCategoryJunction(conn)
	createProduct := productsqlrepo.CreateProduct(conn)
	createStock := stocksqlrepo.CreateStock(conn)
	createSKU := skusqlrepo.CreateSKU(conn)
	createSKUVariant := skuvariantsqlrepo.CreateSKUVariant(conn)
	createSKUPrice := skupricesqlrepo.CreateSKUPrice(conn)
	ee.POST("/skus", CreateSKUs(createSKU, createStock, createProduct, createSKUVariant, createSKUPrice))
	// getCategoryById := categorysqlrepo.GetCategoryByID(conn)
	// getProductCategories := categorysqlrepo.GetProductCategories(conn, getCategoryById)
	getProductByID := productsqlrepo.GetProductByID(conn)
	getProducts := productsqlrepo.GetProducts(conn)
	getSKUs := skusqlrepo.GetSKUs(conn)
	getSKUStockBySKUID := stocksqlrepo.GetSKUStockBySKUID(conn)
	ee.GET("/skus3", GetSKUs(getProducts, getSKUs, getSKUStockBySKUID))

	skuSQLFetcher := skusqlrepo.NewSKUSQLFetcher(conn)

	prod := productsqlrepo.ProductSQLJoin(&skuSQLFetcher, "sku.product_id")
	skuSQLFetcher.Receiver.Product = prod

	st := stocksqlrepo.SKUStockSQLJoin(&skuSQLFetcher)
	skuSQLFetcher.Receiver.Stock = st

	sp := skupricesqlrepo.SKUPriceSQLJoin(&skuSQLFetcher)

	skuSQLFetcher.Receiver.Price = sp

	ee.GET("/skus", GetSKUs3(skuSQLFetcher.Fetch))

	deleteSKUByID := skusqlrepo.DeleteSKUById(conn)
	deleteStockBySKUID := stocksqlrepo.DeleteStockBySKUID(conn)
	deleteSKUOutBySKUID := skuoutsqlrepo.DeleteSKUOutBySKUID(conn)
	ee.DELETE("/skus/:id", DeleteSKUByID(deleteSKUByID, deleteStockBySKUID, deleteSKUOutBySKUID))

	updateSKUByID := skusqlrepo.UpdateSKUByID(conn)

	ee.POST("/skus/:id", UpdateSKUByID(updateSKUByID))

	getRunningLowStocks := stocksqlrepo.GetRunningLowStocks(conn)
	getSKUByID := skusqlrepo.GetSKUByID(conn)
	ee.GET("/skus/lowstocks/", GetRunningLowStocks(getRunningLowStocks, getSKUByID, getProductByID))

	ee.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
}

type createSKUPost struct {
	Product  *product.Product                   `json:"product" validate:"required"`
	Stock    *stock.Stock                       `json:"stock" validate:"required"`
	SKU      *stockkeepingunit.StockKeepingUnit `json:"sku" validate:"required"`
	Price    *skuprice.SKUPrice                 `json:"price" validate:"required"`
	Variants []skuvariant.SKUVariant            `json:"variants"`
}

func validateCreateSKUPost(cskup createSKUPost) error {
	validate := validator.New()
	return validate.Struct(cskup)
}
func CreateSKUs(
	createSKU skurepo.CreateSKUFunc,
	createStock stockrepo.CreateSKUStockFunc,
	createProduct productrepo.CreateProductFunc,
	createSKUVariant skuvariantrepo.CreateSKUVariantFunc,
	createSKUPrice skupricerepo.CrateSKUPriceFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		var posts []createSKUPost
		err := ectx.Bind(&posts)

		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		for _, post := range posts {
			// validating the post request
			err = validateCreateSKUPost(post)
			if err != nil {
				return ectx.JSON(http.StatusBadRequest, err)
			}
			// if given product id is zero then the product haven't been created yet
			// so create the product
			if post.Product.ID == 0 {
				post.Product.CreatedAt = time.Now().Local()
				*post.Product, err = createProduct(ctx, *post.Product)
				if err != nil {
					return ectx.JSON(http.StatusInternalServerError, err)
				}
			}
			post.SKU.ProductID = post.Product.ID
			post.SKU.CreatedAt = time.Now().Local()
			*post.SKU, err = createSKU(ctx, *post.SKU)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}

			post.Stock.SKUID = post.SKU.ID
			post.Price.SKUID = post.SKU.ID

			*post.Stock, err = createStock(ctx, *post.Stock)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			*post.Price, err = createSKUPrice(ctx, *post.Price)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}

			for _, v := range post.Variants {
				v.SKUID = post.SKU.ID
				createSKUVariant(ctx, v)
			}

		}
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusCreated, posts)
	}
}
func GetSKUs2(
	getSKUs skurepo.GetSKUs,
	getProduct productrepo.GetProductByIDFunc,
	getStock stockrepo.GetSKUStockBySKUID,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		skus, err := getSKUs(ctx, nil)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		type skuViewModel struct {
			stockkeepingunit.StockKeepingUnit
			Product product.Product `json:"product"`
			Stock   stock.Stock     `json:"stock"`
		}
		skuProducts := map[int64]*product.Product{}
		for _, sku := range skus {
			skuProducts[sku.ProductID] = nil
		}
		for productId := range skuProducts {
			skuProducts[productId], err = getProduct(ctx, productId)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
		}

		resp := []skuViewModel{}
		for _, sku := range skus {
			st, err := getStock(ctx, sku.ID)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			svm := skuViewModel{
				StockKeepingUnit: sku,
				Product:          *skuProducts[sku.ProductID],
				Stock:            *st,
			}
			resp = append(resp, svm)
		}
		return ectx.JSON(http.StatusOK, resp)
	}
}

func GetSKUs3(
	getSKUs skurepo.GetSKUsFunc,
) echo.HandlerFunc {
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
		keyword := urlQuery.Get("q")
		productNameFilter := options.Filter{
			By:       "product.name",
			Operator: "LIKE",
			Value:    keyword,
		}
		opts := options.Options{Filters: []options.Filter{productNameFilter}}
		skus, err := getSKUs(ctx, &opts)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, models.RESTErrorResponse{Message: err.Error()})
		}
		return ectx.JSON(http.StatusOK, skus)
	}
}

func GetSKUs(
	getProducts productrepo.GetProductsFunc,
	getSKUs skurepo.GetSKUs,
	getSKUStockBySKUID stockrepo.GetSKUStockBySKUID,
) echo.HandlerFunc {
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
		var opts *options.Options
		name := ectx.Request().URL.Query().Get("name")
		if name != "" {
			opts = &options.Options{}
			ft := options.Filter{
				By:       "name",
				Value:    name,
				Operator: "LIKE",
			}
			opts.Filters = []options.Filter{ft}
		}
		products, err := getProducts(ctx, opts)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		completeSKUs := []stockkeepingunit.CompleteSKU{}
		for _, p := range products {
			skuOpts := options.Options{}
			skuProductFilter := options.Filter{
				By:       "product_id",
				Value:    p.ID,
				Operator: "=",
			}
			skuOpts.Filters = []options.Filter{skuProductFilter}
			skus, err := getSKUs(ctx, &skuOpts)
			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			for _, sku := range skus {
				st, err := getSKUStockBySKUID(ctx, sku.ID)
				if err != nil {
					return ectx.JSON(http.StatusInternalServerError, err)
				}
				cSKU := stockkeepingunit.CompleteSKU{
					StockKeepingUnit: sku,
					Product:          p,
					Stock:            *st,
				}
				completeSKUs = append(completeSKUs, cSKU)
			}
		}
		return ectx.JSON(http.StatusOK, completeSKUs)
	}
}

func DeleteSKUByID(
	deleteSKUByID skurepo.DeleteSKUByID,
	deleteStockBySKUID stockrepo.DeleteStockBySKUID,
	deleteSKUOutBySKUID skuoutrepo.DeleteSKUOutBySKUIDFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		err = deleteSKUByID(ctx, id)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		err = deleteStockBySKUID(ctx, id)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		err = deleteSKUOutBySKUID(ctx, id)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}

		return ectx.JSON(http.StatusOK, "Succeded deleting sku")
	}
}
func UpdateSKUByID(updateSKU skurepo.UpdateSKUFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		var post struct {
			SKU stockkeepingunit.StockKeepingUnit `json:"SKU"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}

		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		post.SKU.ID = id

		sku, err := updateSKU(ctx, post.SKU)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, sku)
	}
}

func GetRunningLowStocks(
	getRunningLowStocks stockrepo.GetRunnigLowStocksFunc,
	getSKUByID skurepo.GetSKUByIDFunc,
	getProductByID productrepo.GetProductByIDFunc,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		log.Println("it's confirmed")
		ctx := ectx.Request().Context()
		if ctx == nil {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
		}
		stocks, err := getRunningLowStocks(ctx)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		type respModel struct {
			stockkeepingunit.StockKeepingUnit
			Product product.Product `json:"product"`
			Stock   stock.Stock     `json:"stock"`
		}
		resp := []respModel{}
		for _, st := range stocks {
			sku, err := getSKUByID(ctx, st.SKUID)
			log.Println(err, st)
			if err != nil {
				continue
			}

			prod, err := getProductByID(ctx, sku.ProductID)

			if err != nil {
				return ectx.JSON(http.StatusInternalServerError, err)
			}
			r := respModel{
				Stock:            st,
				StockKeepingUnit: *sku,
				Product:          *prod,
			}

			resp = append(resp, r)
		}
		return ectx.JSON(http.StatusOK, resp)
	}
}
