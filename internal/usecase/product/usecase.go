package product

import (
	"context"
	mProduct "github/fadlinux/edot/internal/model/product"
)

type (
	// Param : param struct that sent to usecase
	Param struct {
		ID       int64
		Name     string
		Phone    string
		Email    string
		Password string
		Page     int64
		Size     int64
	}
)

// Usecase : contract of functions in Usecase layer
type Usecase interface {
	FetchProduct(ctx context.Context, params mProduct.SearchParam) (result mProduct.Response, err error)
	AddProduct(ctx context.Context, params mProduct.Product) (lastID int64, err error)
}
