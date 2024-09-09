package mysql

import (
	"database/sql"

	rOrder "github/fadlinux/edot/internal/repository/order"
)

type (
	mysqlOrderRepo struct {
		MySqlConnection *sql.DB
	}
)

// NewMySqlOrderRepo create new object for Order Postgre Repository
func NewMySqlOrderRepo(Conn *sql.DB) rOrder.Repository {
	repo := &mysqlOrderRepo{Conn}

	// initialize preparestatements
	repo.prepareAddOrderStmt()
	repo.prepareAddOrderDetailStmt()
	repo.prepareAddReserverStockStmt()

	return repo
}
