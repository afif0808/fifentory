package productsqlrepo

import (
	"context"
	"fifentory/category"
	"fifentory/product"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/require"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetProductById(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	rows := mock.NewRows(strings.Split(productFields, ",")).AddRow(1, "Shirt", time.Now())
	mock.ExpectQuery(strings.Replace(getProductByIdQuery, "?", "\\?", -1)).WillReturnRows(rows)
	mockCategories := []category.Category{}
	err = faker.FakeData(&mockCategories)
	prod, err := GetProductByID(db)(context.Background(), 1)
	require.NoError(t, mock.ExpectationsWereMet())

	assert.NoError(t, err)
	assert.NotEqual(t, (product.Product{}), prod)
}
