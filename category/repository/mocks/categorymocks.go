package categoryrepomocks

import (
	"context"
	"fifentory/category"
	categoryrepo "fifentory/category/repository"
)

func MockGetProductCategoriesFunc(ret ...interface{}) categoryrepo.GetProductCategoriesFunc {
	return func(ctx context.Context, productID int64) ([]category.Category, error) {
		var categories []category.Category
		var err error
		if ret0, ok := ret[0].([]category.Category); ok {
			categories = ret0
		}
		if ret1, ok := ret[1].(error); ok {
			err = ret1
		}
		return categories, err
	}
}
