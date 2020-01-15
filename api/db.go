package api

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DBconn connection for database
func DBconn() (db *sql.DB) {

	dbDriver := os.Getenv("dbDriver")
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")
	dbName := os.Getenv("dbName")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
