package skuresthandler

import (
	"context"
	"database/sql"
	"fifentory/options"
	"fifentory/product"
	productrepo "fifentory/product/repository"
	productsqlrepo "fifentory/product/repository/sql"
	"fifentory/stock"
	stockrepo "fifentory/stock/repository"
	stocksqlrepo "fifentory/stock/repository/sql"
	"fifentory/stockkeepingunit"
	skusqlrepo "fifentory/stockkeepingunit/repository/sql"
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
	ee.POST("/skus", CreateSKUs(createSKU, createStock, createProduct))
	// getCategoryById := categorysqlrepo.GetCategoryByID(conn)
	// getProductCategories := categorysqlrepo.GetProductCategories(conn, getCategoryById)
	// getProductById := productsqlrepo.GetProductByID(conn)
	getProducts := productsqlrepo.GetProducts(conn)
	getSKUs := skusqlrepo.GetSKUs(conn)
	getSKUStockBySKUID := stocksqlrepo.GetSKUStockBySKUID(conn)
	ee.GET("/skus", GetSKUs(getProducts, getSKUs, getSKUStockBySKUID))

	deleteSKUByID := skusqlrepo.DeleteSKUById(conn)
	ee.DELETE("/skus/:id", DeleteSKUByID(deleteSKUByID))

	updateSKUByID := skusqlrepo.UpdateSKUByID(conn)

	ee.POST("/skus/:id", UpdateSKUByID(updateSKUByID))

	ee.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
}

type createSKUPost struct {
	Product *product.Product                   `json:"product" validate:"required"`
	Stock   *stock.Stock                       `json:"stock" validate:"required"`
	SKU     *stockkeepingunit.StockKeepingUnit `json:"sku" validate:"required"`
}

func validateCreateSKUPost(cskup createSKUPost) error {
	validate := validator.New()
	return validate.Struct(cskup)
}
func CreateSKUs(
	createSKU skurepo.CreateSKUFunc,
	createStock stockrepo.CreateSKUStockFunc,
	createProduct productrepo.CreateProductFunc,
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
			*post.Stock, err = createStock(ctx, *post.Stock)
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

func GetSKUs(
	getProducts productrepo.GetProductsFunc,
	getSKUs skurepo.GetSKUs,
	getSKUStockBySKUID stockrepo.GetSKUStockBySKUID,
) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		ctx := ectx.Request().Context()
		if ctx == nil {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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

func DeleteSKUByID(deleteSKU skurepo.DeleteSKUByID) echo.HandlerFunc {
	return func(ectx echo.Context) error {

		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		err = deleteSKU(ctx, id)
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
