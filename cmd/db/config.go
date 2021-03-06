package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
const (
		host     = "db"
		port     = 5432
		user     = "gohub"
		password = "gohub_dev_password"
		dbname   = "gohub_dev"
		)

// NOTE: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
func ConnectionDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
