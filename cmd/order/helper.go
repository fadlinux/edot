package order

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github/fadlinux/edot/common/util/config"
	log "github/fadlinux/edot/common/util/log"

	_ "github.com/go-sql-driver/mysql"
)

func newDBConnection(driver, host string) (conn *sql.DB) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatalln("Failed to init mysql connection", err)
	} else {
		conn.SetConnMaxLifetime(time.Duration(config.GetInt("mysql.db_conn_max_lifetime")) * time.Second)
	}
	return
}
