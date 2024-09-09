package product

import (
	rProduct "github/fadlinux/edot/internal/repository/product"
)

type (
	productUC struct {
		mysqlProductRepo rProduct.Repository
	}
)

// NewProductUC : create new instance for Product usecase
func NewProductUC(mysqlProduct rProduct.Repository) Usecase {
	return &productUC{
		mysqlProductRepo: mysqlProduct,
	}
}
