package shop

import (
	rShop "github/fadlinux/edot/internal/repository/shop"
)

type (
	shopUC struct {
		mysqlShopRepo rShop.Repository
	}
)

// NewShopUC : create new instance for Shop usecase
func NewShopUC(mysqlShop rShop.Repository) Usecase {
	return &shopUC{
		mysqlShopRepo: mysqlShop,
	}
}
