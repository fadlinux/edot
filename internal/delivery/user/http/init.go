package http

import (
	uUser "github/fadlinux/edot/internal/usecase/user"
)

// Delivery : deliver response using http
type Delivery struct {
	userUC uUser.Usecase
}

// NewUserHTTP : create new instance for User http delivery
func NewUserHTTP(userUC uUser.Usecase) Delivery {
	return Delivery{userUC}
}
