package mysql

import (
	"database/sql"

	log "github/fadlinux/edot/common/util/log"
	//Postgressql driver
)

var (
	fetchUserPrepareStmt      *sql.Stmt
	countUserPrepareStmt      *sql.Stmt
	addUserPrepareStmt        *sql.Stmt
	countUserLoginPrepareStmt *sql.Stmt
)

func (a *mysqlUserRepo) prepareFetchUserStmt() {
	var err error
	fetchUserPrepareStmt, err = a.MySqlConnection.Prepare("SELECT id, name, email, phone, created_at, updated_at,password_hash FROM users WHERE email = ? OR phone = ? ")
	if err != nil {
		log.Fatal("[User Repo] Prepare select statement fail :", err)
	}
}

func (a *mysqlUserRepo) prepareCountUserStmt() {
	var err error
	countUserPrepareStmt, err = a.MySqlConnection.Prepare("SELECT COUNT(id) FROM users WHERE name LIKE ?")
	if err != nil {
		log.Fatal("[User Repo] Prepare count statement fail :", err)
	}
}

func (a *mysqlUserRepo) prepareCountUserLoginPrepareStmt() {
	var err error
	countUserLoginPrepareStmt, err = a.MySqlConnection.Prepare("SELECT COUNT(id) FROM users WHERE phone = ? OR email = ?")
	if err != nil {
		log.Fatal("[User Repo] Prepare count statement fail :", err)
	}
}

func (a *mysqlUserRepo) prepareAddUserStmt() {
	var err error
	addUserPrepareStmt, err = a.MySqlConnection.Prepare("INSERT INTO users (name, phone, email, password_hash) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal("[User Repo] Prepare add statement fail :", err)
	}
}
