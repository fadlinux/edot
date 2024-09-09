package order

import (
	rOrder "github/fadlinux/edot/internal/repository/order"
	rProduct "github/fadlinux/edot/internal/repository/product"
)

type (
	orderUC struct {
		mysqlOrderRepo   rOrder.Repository
		mysqlProductRepo rProduct.Repository
	}
)

// NewOrderUC : create new instance for Order usecase
func NewOrderUC(mysqlOrder rOrder.Repository, mysqlProduct rProduct.Repository) Usecase {
	return &orderUC{
		mysqlOrderRepo:   mysqlOrder,
		mysqlProductRepo: mysqlProduct,
	}
}
