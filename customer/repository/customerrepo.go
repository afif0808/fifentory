package customerrepo

import (
	"context"
	"fifentory/customer"
)

type GetCustomersFunc func(context.Context) ([]customer.Customer, error)

type GetCustomerByIDFunc func(ctx context.Context, id int64) (*customer.Customer, error)
