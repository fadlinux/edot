package mysql

import (
	"database/sql"

	rUser "github/fadlinux/edot/internal/repository/user"
)

type (
	mysqlUserRepo struct {
		MySqlConnection *sql.DB
	}
)

// NewMySqlUserRepo create new object for User Postgre Repository
func NewMySqlUserRepo(Conn *sql.DB) rUser.Repository {
	repo := &mysqlUserRepo{Conn}

	// initialize preparestatements
	repo.prepareFetchUserStmt()
	repo.prepareCountUserStmt()
	repo.prepareAddUserStmt()
	repo.prepareCountUserLoginPrepareStmt()

	return repo
}
