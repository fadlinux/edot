package shop

import (
	"context"
	mShop "github/fadlinux/edot/internal/model/shop"
	"log"

	rShop "github/fadlinux/edot/internal/repository/shop"
)

func (a *shopUC) FetchShop(ctx context.Context, params mShop.Shop) (result mShop.Response, err error) {
	shopData, err := a.mysqlShopRepo.FetchShopByPhoneEmail(ctx, rShop.Param{
		Name: params.Name,
	})

	if err != nil {
		log.Println("[Shop Repository] Fail to fetch data", err)
		return
	}

	for _, v := range shopData {
		result.Data = append(result.Data, mShop.Data{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return
}

func (a *shopUC) AddShop(ctx context.Context, params mShop.Shop) (lastID int64, err error) {
	lastID, err = a.mysqlShopRepo.AddShop(ctx, rShop.Param{
		Name: params.Name,
	})
	return
}

func (a *shopUC) GetShopLogin(ctx context.Context, params mShop.Shop) (totalData int64, err error) {
	totalData, err = a.mysqlShopRepo.CountShopByPhoneEmail(ctx, rShop.Param{
		Email: params.Name,
	})

	return
}
