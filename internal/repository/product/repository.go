package product

import (
	"context"
	mProduct "github/fadlinux/edot/internal/model/product"
)

type (
	// Param : param struct that sent to repository
	Param struct {
		ID          int64
		Name        string
		Description string
		Price       int
		Stock       int
		WarehouseId string
		Page        int64
		Size        int64
		Q           string
	}
)

// Repository : contract of functions in repository layer
type Repository interface {
	AddProduct(ctx context.Context, params Param) (lastID int64, err error)
	FetchSearchProduct(ctx context.Context, params mProduct.SearchParam) (result []mProduct.Product, err error)
}
