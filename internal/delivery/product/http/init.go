package http

import (
	uProduct "github/fadlinux/edot/internal/usecase/product"
)

// Delivery : deliver response using http
type Delivery struct {
	productUC uProduct.Usecase
}

// NewProductHTTP : create new instance for Product http delivery
func NewProductHTTP(productUC uProduct.Usecase) Delivery {
	return Delivery{productUC}
}
