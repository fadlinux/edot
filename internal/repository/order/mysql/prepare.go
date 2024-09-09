package mysql

import (
	"database/sql"

	log "github/fadlinux/edot/common/util/log"
)

var (
	addOrderPrepareStmt         *sql.Stmt
	addOrderDetailPrepareStmt   *sql.Stmt
	addReserverStockPrepareStmt *sql.Stmt
)

func (a *mysqlOrderRepo) prepareAddOrderStmt() {
	var err error
	addOrderPrepareStmt, err = a.MySqlConnection.Prepare("INSERT INTO orders (user_id, status, created_at ) VALUES (?,?,?)")
	if err != nil {
		log.Fatal("[Order Repo] Prepare add statement fail :", err)
	}
}

func (a *mysqlOrderRepo) prepareAddOrderDetailStmt() {
	var err error
	addOrderDetailPrepareStmt, err = a.MySqlConnection.Prepare("INSERT INTO order_items (order_id,product_id,quantity,price ) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal("[Order Repo] Prepare add statement fail :", err)
	}
}

func (a *mysqlOrderRepo) prepareAddReserverStockStmt() {
	var err error
	addReserverStockPrepareStmt, err = a.MySqlConnection.Prepare("INSERT INTO reserved_stock (order_id,product_id,reserved_quantity, expires_at ) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal("[Order Repo] Prepare add statement fail :", err)
	}
}
