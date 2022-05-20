package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
const (
		host     = "db"
		port     = 5432
		user     = "postgres"
		password = "gohub_dev_password"
		dbname   = "gohub_dev"
		)

type User struct {
	Id   int `db:"id"`
	LastName string `db:"last_name"`
	FirstName string `db:"first_name"`
	Age int `db:"age"`
}
// NOTE: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
func ConnectionDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from gohub_user")
	if err != nil {
		log.Println(err)
	}

	var u User
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("id: %d, username: %s %s, age: %d\n", u.Id, u.LastName, u.FirstName, u.Age)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}


	fmt.Println("Successfully connected!")
}
