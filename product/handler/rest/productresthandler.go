package productresthandler

import (
	"context"
	"database/sql"
	"fifentory/category"
	categoryrepo "fifentory/category/repository"
	categorysqlrepo "fifentory/category/repository/sql"
	"fifentory/models"
	"fifentory/product"
	productrepo "fifentory/product/repository"
	productsqlrepo "fifentory/product/repository/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func InjectProductRESTHandler(conn *sql.DB, ee *echo.Echo) {
	getCategoryById := categorysqlrepo.GetCategoryByID(conn)
	getProductCategories := categorysqlrepo.GetProductCategories(conn, getCategoryById)
	getProductByID := productsqlrepo.GetProductByID(conn)
	ee.GET("/products/:id", GetProductByID(getProductByID, getProductCategories))

	createProductCategoryJunction := categorysqlrepo.CreateProductCategoryJunction(conn)
	createProduct := productsqlrepo.CreateProduct(conn)
	ee.POST("/products", CreateProduct(createProduct, createProductCategoryJunction))

	updateProductById := productsqlrepo.UpdateProductById(conn)
	ee.POST("/products/:id", UpdateProductById(updateProductById))

}

type ErrorResponse struct {
	Msg error `json:"message"`
}

func GetProductByID(
	getProduct productrepo.GetProductByIDFunc,
	getProductCategories categoryrepo.GetProductCategoriesFunc,
) echo.HandlerFunc {
	return echo.HandlerFunc(func(ectx echo.Context) error {
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusNotFound, ErrorResponse{Msg: models.ErrNotFound})
		}
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		prod, err := getProduct(ctx, id)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, ErrorResponse{Msg: models.ErrInternalServerError})

		}

		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, ErrorResponse{Msg: models.ErrInternalServerError})
		}

		categories, err := getProductCategories(ctx, prod.ID)

		var resp struct {
			product.Product
			Categories []category.Category `json:"categories,omitempty"`
		}
		resp.Product = *prod
		resp.Categories = categories
		return ectx.JSON(http.StatusOK, resp)
	})
}

func CreateProduct(
	createProduct productrepo.CreateProductFunc,
	createProductCategories categoryrepo.CreateProductCategoryJunctionFunc,
) echo.HandlerFunc {
	return echo.HandlerFunc(func(ectx echo.Context) error {
		var post struct {
			Product    product.Product     `json:"product"`
			Categories []category.Category `json:"categories,omitempty"`
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, ErrorResponse{Msg: err})
		}

		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		if post.Product.CreatedAt == (time.Time{}) {
			post.Product.CreatedAt = time.Now().Local()
		}
		post.Product, err = createProduct(ctx, post.Product)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, ErrorResponse{Msg: err})
		}
		for _, c := range post.Categories {
			_, err = createProductCategories(
				ctx,
				category.ProductCategoryJunction{ProductID: post.Product.ID, CategoryID: c.ID},
			)
		}
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusCreated, post)
	})
}

func UpdateProductById(updateProduct productrepo.UpdateProductFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		var post struct {
			Product product.Product `json:"product"`
		}
		ctx := ectx.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		err := ectx.Bind(&post)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 64)
		if err != nil {
			return ectx.JSON(http.StatusBadRequest, err)
		}
		post.Product.ID = id
		prod, err := updateProduct(ctx, post.Product)
		if err != nil {
			return ectx.JSON(http.StatusInternalServerError, err)
		}
		return ectx.JSON(http.StatusOK, prod)
	}
}
