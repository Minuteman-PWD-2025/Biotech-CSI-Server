package main

import (
	fmt "fmt"
	"strings"

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
	//we need to fix this one to match later ones and use a for loop instead of messy formatting

	rows, err := db.Query("INSERT INTO " + WhichTable + " " + cols + "\nVALUES " + Data)
	if err != nil {
		panic(err)
	}
	return *rows
}
func AlterThing(WhichTable string, allUpdates []string, allWheres []string) {
	for i := 0; i < len(allUpdates); i++ {
		splitInTwain := strings.Split(allUpdates[i], ",")
		for i := 0; i < len(allWheres); i++ {
			whereInTwain := strings.Split(allWheres[i], ",")
			qStr := "UPDATE " + WhichTable + "\nSET " + splitInTwain[0] + "=" + splitInTwain[1] + "\nWHERE " + whereInTwain[0] + whereInTwain[1] + ";"
			fmt.Println(qStr)
			db.Query(qStr)
		}
	}

}
func DeleteRow(WhichTable string, allWheres []string) {
	for i := 0; i < len(allWheres); i++ {
		splitInTwain := strings.Split(allWheres[i], ",")
		qStr := "DELETE FROM " + WhichTable + "\nWHERE " + splitInTwain[0] + splitInTwain[1] + ";"
		fmt.Print(qStr)
		db.Query(qStr)
	}
}
