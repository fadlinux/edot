package shop

import (
	"context"
	mShop "github/fadlinux/edot/internal/model/shop"
)

type (
	// Param : param struct that sent to repository
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

// Repository : contract of functions in repository layer
type Repository interface {
	AddShop(ctx context.Context, params Param) (lastID int64, err error)
	FetchShopByPhoneEmail(ctx context.Context, params Param) (result []mShop.Shop, err error)
	CountShopByPhoneEmail(ctx context.Context, params Param) (total int64, err error)
}
