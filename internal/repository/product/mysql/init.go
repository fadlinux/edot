package mysql

import (
	"database/sql"

	rProduct "github/fadlinux/edot/internal/repository/product"
)

type (
	mysqlProductRepo struct {
		MySqlConnection *sql.DB
	}
)

// NewMySqlProductRepo create new object for Product Postgre Repository
func NewMySqlProductRepo(Conn *sql.DB) rProduct.Repository {
	repo := &mysqlProductRepo{Conn}

	// initialize preparestatements
	repo.prepareFetchProductStmt()
	repo.prepareAddProductStmt()
	repo.prepareCountFetchProductStmt()

	return repo
}
