package main

import (
	fmt "fmt"
	"strconv"

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

var db *sql.DB

func EnableServer() {
	mydb := ConnectToTable()
	db = mydb
}

//func testConnection() (error, *sql.DB) {
// fmt.Printf("Testing Connection to Database...\n")
// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_pass, db_name)
// db, err := sql.Open("postgres", connStr)
// if err != nil {
// 	return err, nil
// }

//defer db.Close()

// rows, err := db.Query("SELECT * FROM people;")

// if err != nil {
// 	panic(err)
// }

// var id string
// var name string
// for rows.Next() {
// 	rows.Scan(&id, &name)
// 	fmt.Printf("ID: %s\nName: %s\n\n", id, name)
// }

//return err, db
//}
func ConnectToTable() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_pass, db_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
func GetTable(WhichTable string) sql.Rows {

	rows, err := db.Query("SELECT * FROM " + WhichTable + ";")
	if err != nil {
		panic(err)
	}
	return *rows
}
func AddNew(WhichTable string, cols string, Data string) sql.Rows {
	//validate that it doesnt already exist

	rows, err := db.Query("INSERT INTO " + WhichTable + " " + cols + "\nVALUES " + Data)
	if err != nil {
		panic(err)
	}
	return *rows
}
func GetId(WhichTable string) int {

	validRow := false
	i := 0
	for {
		i++

		checkRows, err := db.Query("SELECT user_id FROM " + WhichTable + " t WHERE t.user_id = " + strconv.Itoa(i))
		if err != nil {
			panic(err)
		}
		cols, err := checkRows.Columns()
		if len(cols) == 0 {
			validRow = true
		}
		if validRow {
			break
		}
	}
	return i
}
