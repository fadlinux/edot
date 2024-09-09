package user

import (
	"database/sql"
	"time"

	"github/fadlinux/edot/common/util/config"
	log "github/fadlinux/edot/common/util/log"

	_ "github.com/go-sql-driver/mysql"
)

func newDBConnection(driver, host string) (conn *sql.DB) {
	conn, err := sql.Open(driver, host)
	if err != nil {
		log.Fatalln("Failed to init mysql connection", err)
	} else {
		conn.SetConnMaxLifetime(time.Duration(config.GetInt("mysql.db_conn_max_lifetime")) * time.Second)
	}
	return
}
