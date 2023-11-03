package main

import (
	fmt "fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	db_name = "postgres"
	db_user = "postgres"
	db_pass = "admin"
	db_host = "127.0.0.1"
	db_port = "5432"
)

func testConnection() error {
	fmt.Printf("Testing Connection to Database...\n")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_pass, db_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	defer db.Close()

	var rows []string

	query, err := db.Query("SELECT * FROM people;")

	if err != nil {
		panic(err)
	}

	query.Scan(&rows)

	fmt.Printf("Found rows: %s \n", rows)

	return err
}
