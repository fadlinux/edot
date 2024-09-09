package mysql

import (
	"context"
	"log"
	"time"

	mOrder "github/fadlinux/edot/internal/model/order"
)

func (a *mysqlOrderRepo) AddOrder(ctx context.Context, params mOrder.Order) (lastID int64, err error) {
	var datetime = time.Now()
	dt := datetime.Format("2006-01-02 15:04:05")

	res, err := addOrderPrepareStmt.Exec(params.UserID, "active", dt)
	if err != nil {
		log.Println("error : addOrderPrepareStmt.QueryRow :", err.Error())
		lastID = 0
		return
	}
	lastID, _ = res.LastInsertId()

	return lastID, err
}

func (a *mysqlOrderRepo) AddOrderDetail(ctx context.Context, orderId int64, data mOrder.Order) (err error) {
	for _, y := range data.Products {
		_ = addOrderDetailPrepareStmt.QueryRow(orderId, y.ProductID, y.Quantity, y.Price)
	}

	return
}

func (a *mysqlOrderRepo) AddReserveStock(ctx context.Context, orderId int64, params mOrder.Order) (err error) {
	var datetime = time.Now()
	datetime.Add(time.Minute * 30)
	dt := datetime.Format("2006-01-02 15:04:05")

	for _, y := range params.Products {
		_ = addReserverStockPrepareStmt.QueryRow(orderId, y.ProductID, y.Quantity, dt).Scan()
	}

	return
}
