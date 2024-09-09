package postgre

import (
	"database/sql"
	"log"
	"time"
)

// NewDBConnection : create new connection for Database
func NewDBConnection(driver, host string) (conn *sql.DB) {
	conn, err := sql.Open(driver, host)
	if err != nil {
		log.Fatalln("Failed to init postgre", err)
	} else {
		conn.SetConnMaxLifetime(2 * time.Second)
	}
	return
}
