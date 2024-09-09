package mysql

import (
	"database/sql"

	log "github/fadlinux/edot/common/util/log"
	//Postgressql driver
)

var (
	fetchShopPrepareStmt      *sql.Stmt
	countShopPrepareStmt      *sql.Stmt
	addShopPrepareStmt        *sql.Stmt
	countShopLoginPrepareStmt *sql.Stmt
)

func (a *mysqlShopRepo) prepareFetchShopStmt() {
	var err error
	fetchShopPrepareStmt, err = a.MySqlConnection.Prepare("SELECT id, name,  updated_at FROM shops WHERE name = ? ")
	if err != nil {
		log.Fatal("[Shop Repo] Prepare select statement fail :", err)
	}
}

func (a *mysqlShopRepo) prepareCountShopStmt() {
	var err error
	countShopPrepareStmt, err = a.MySqlConnection.Prepare("SELECT COUNT(id) FROM shops WHERE name LIKE ?")
	if err != nil {
		log.Fatal("[Shop Repo] Prepare count statement fail :", err)
	}
}

func (a *mysqlShopRepo) prepareAddShopStmt() {
	var err error
	addShopPrepareStmt, err = a.MySqlConnection.Prepare("INSERT INTO shops (name,created_at) VALUES (?,?)")
	if err != nil {
		log.Fatal("[Shop Repo] Prepare add statement fail :", err)
	}
}
