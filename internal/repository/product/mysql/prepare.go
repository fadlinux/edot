package mysql

import (
	"database/sql"

	log "github/fadlinux/edot/common/util/log"
	//Postgressql driver
)

var (
	fetchProductPrepareStmt      *sql.Stmt
	addProductPrepareStmt        *sql.Stmt
	fetchCountProductPrepareStmt *sql.Stmt
)

func (a *mysqlProductRepo) prepareAddProductStmt() {
	var err error
	addProductPrepareStmt, err = a.MySqlConnection.Prepare("INSERT INTO products (name, description, price, stock, warehouse_id, created_at ) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Fatal("[Product Repo] Prepare add statement fail :", err)
	}
}

func (a *mysqlProductRepo) prepareFetchProductStmt() {
	var err error
	fetchProductPrepareStmt, err = a.MySqlConnection.Prepare("SELECT id, name, description, price, stock, warehouse_id, created_at, updated_at FROM products WHERE stock > 0 and name LIKE ? ORDER BY id desc LIMIT ?,? ")
	if err != nil {
		log.Fatal("[Product Repo] Prepare select statement fail :", err)
	}
}

func (a *mysqlProductRepo) prepareCountFetchProductStmt() {
	var err error
	fetchCountProductPrepareStmt, err = a.MySqlConnection.Prepare("SELECT COUNT(id) FROM products WHERE stock > 0 AND name LIKE ?")
	if err != nil {
		log.Fatal("[Product Repo] Prepare select count statement fail :", err)
	}
}
