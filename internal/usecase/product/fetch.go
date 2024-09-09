package product

import (
	"context"
	mProduct "github/fadlinux/edot/internal/model/product"
	"log"

	rProduct "github/fadlinux/edot/internal/repository/product"
)

func (a *productUC) FetchProduct(ctx context.Context, params mProduct.SearchParam) (result mProduct.Response, err error) {
	productData, err := a.mysqlProductRepo.FetchSearchProduct(ctx, mProduct.SearchParam{
		Q:    params.Q,
		Size: params.Size,
		Page: params.Page,
	})

	if err != nil {
		log.Println("[Product Repository] Fail to fetch data", err)
		return
	}

	for _, v := range productData {
		result.Data = append(result.Data, mProduct.Data{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			WarehouseId: v.WarehouseId,
			Stock:       v.Stock,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return
}

func (a *productUC) AddProduct(ctx context.Context, params mProduct.Product) (lastID int64, err error) {
	lastID, err = a.mysqlProductRepo.AddProduct(ctx, rProduct.Param{
		Name:        params.Name,
		Description: params.Description,
		Price:       int(params.Price),
		WarehouseId: params.WarehouseId,
		Stock:       params.Stock,
	})
	return
}
