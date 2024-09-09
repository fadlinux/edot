package http

import (
	uShop "github/fadlinux/edot/internal/usecase/shop"
)

// Delivery : deliver response using http
type Delivery struct {
	shopUC uShop.Usecase
}

// NewShopHTTP : create new instance for Shop http delivery
func NewShopHTTP(shopUC uShop.Usecase) Delivery {
	return Delivery{shopUC}
}
