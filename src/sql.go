package main

import (
	fmt "fmt"
	os "os"

	"database/sql"

	_ "github.com/lib/pq"
)

var db_name = os.Getenv("db_name")
var db_user = os.Getenv("db_user")

func testConnection() error {
	fmt.Printf("Testing Connection to Database...\n")
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", db_user, db_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	var rows []string

	query, err := db.Query("SELECT name FROM people;")

	if err != nil {
		panic(err)
	}

	query.Scan(&rows)

	fmt.Printf("Found rows: %s \n", rows)

	defer db.Close()
	return err
}
