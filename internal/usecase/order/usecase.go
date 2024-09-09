package order

import (
	"context"
	mOrder "github/fadlinux/edot/internal/model/order"
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
	FetchOrder(ctx context.Context, params mOrder.SearchParam) (result mOrder.Response, err error)
	AddOrder(ctx context.Context, params mOrder.Order) (lastID int64, err error)
}
