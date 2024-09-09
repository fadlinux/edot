package order

import (
	"context"
	mOrder "github/fadlinux/edot/internal/model/order"
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
	AddOrder(ctx context.Context, params mOrder.Order) (lastID int64, err error)
	AddOrderDetail(ctx context.Context, orderId int64, data mOrder.Order) (err error)
	AddReserveStock(ctx context.Context, orderId int64, params mOrder.Order) (err error)
}
