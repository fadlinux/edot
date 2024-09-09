package http

import (
	uOrder "github/fadlinux/edot/internal/usecase/order"
)

// Delivery : deliver response using http
type Delivery struct {
	orderUC uOrder.Usecase
}

// NewOrderHTTP : create new instance for Order http delivery
func NewOrderHTTP(orderUC uOrder.Usecase) Delivery {
	return Delivery{orderUC}
}
