package shop

import (
	"context"
	mShop "github/fadlinux/edot/internal/model/shop"
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
	FetchShop(ctx context.Context, params mShop.Shop) (result mShop.Response, err error)
	AddShop(ctx context.Context, params mShop.Shop) (lastID int64, err error)
	GetShopLogin(ctx context.Context, params mShop.Shop) (totalData int64, err error)
}
