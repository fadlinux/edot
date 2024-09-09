package order

import (
	"context"
	mOrder "github/fadlinux/edot/internal/model/order"
	"log"
)

func (a *orderUC) AddOrder(ctx context.Context, params mOrder.Order) (lastID int64, err error) {
	lastID, err = a.mysqlOrderRepo.AddOrder(ctx, params)
	if err != nil {
		log.Println("error AddOrder : " + err.Error())
	}

	err = a.mysqlOrderRepo.AddOrderDetail(ctx, lastID, params)
	if err != nil {
		log.Println("error AddOrderDetail : " + err.Error())
	}

	err = a.mysqlOrderRepo.AddReserveStock(ctx, lastID, params)
	if err != nil {
		log.Println("error AddReserveStock : " + err.Error())
	}

	return lastID, err
}
