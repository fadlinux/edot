package mysql

import (
	"database/sql"

	rShop "github/fadlinux/edot/internal/repository/shop"
)

type (
	mysqlShopRepo struct {
		MySqlConnection *sql.DB
	}
)

// NewMySqlShopRepo create new object for Shop Postgre Repository
func NewMySqlShopRepo(Conn *sql.DB) rShop.Repository {
	repo := &mysqlShopRepo{Conn}

	// initialize preparestatements
	repo.prepareFetchShopStmt()
	repo.prepareCountShopStmt()
	repo.prepareAddShopStmt()

	return repo
}
