package main

import (
	"encoding/json"
	"errors"
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
	sep     = "|"
)

var db *sql.DB

func EnableServer() {
	mydb, err := ConnectToTable()
	if err != nil {
		log(true, err.Error())
		//os.Exit(404)
		return
	}
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

// return err, db
// }
func ConnectToTable() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_pass, db_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.New("error opening database connection: " + err.Error())
	}

	// check if the database is accessible
	if err = db.Ping(); err != nil {
		return nil, errors.New("error connecting to the database: " + err.Error())
	}

	return db, nil
}
func GetTable(table string) sql.Rows {

	rows, err := db.Query("SELECT * FROM " + RemSpaces(table) + ";")
	if err != nil {
		panic(err)
	}
	return *rows
}

//Converts an input to the first thing, removing stuff after spaces, to try and prevent SQL injections
func RemSpaces(query string) string {
	splitString := strings.Split(query, " ")
	return splitString[0]
}

func AddNew(table string, cols string, Data string) (sql.Rows, error) {
	//we need to fix this one to match later ones and use a for loop instead of messy formatting

	rows, err := db.Query("INSERT INTO " + table + " " + cols + "\nVALUES " + Data)
	if err != nil {
		err = errors.New("error inserting data: " + err.Error())
	}
	return *rows, err
}
func AlterThing(WhichTable string, allUpdates []string, allWheres []string) error {
	for i := 0; i < len(allUpdates); i++ {
		splitInTwain := strings.Split(allUpdates[i], ",")
		for j := 0; j < len(allWheres); j++ {
			whereInTwain := strings.Split(allWheres[j], ",")
			qStr := "UPDATE " + WhichTable + "\nSET " + splitInTwain[0] + "=" + splitInTwain[1] + "\nWHERE " + whereInTwain[0] + whereInTwain[1] + ";"
			fmt.Println(qStr)
			_, err := db.Query(qStr)
			if err != nil {
				return errors.New("error altering thing: " + err.Error())
			}
		}
	}
	return nil
}
func DeleteRow(WhichTable string, allWheres []string) error {
	for i := 0; i < len(allWheres); i++ {
		splitInTwain := strings.Split(allWheres[i], ",")
		qStr := "DELETE FROM " + WhichTable + "\nWHERE " + splitInTwain[0] + splitInTwain[1] + ";"
		fmt.Print(qStr)
		_, err := db.Query(qStr)
		if err != nil {
			return errors.New("error deleting data: " + err.Error())
		}
	}
	return nil
}
func FormatTableToJSON(table string) []byte {
	rows := GetTable(table)
	cols, _ := rows.Columns()
	leng := len(cols)
	datas := make([]any, leng) // array of references
	fmt.Print(rows.Columns())

	for i := 0; i < leng; i++ {
		datas[i] = new(any)
	}

	var finData []map[string]any

	for rows.Next() {
		tempData := map[string]any{}
		rows.Scan(datas...) // unwrap array of references and pass through

		for i, data := range datas {
			tempData[cols[i]] = (*data.(*any))

		}

		finData = append(finData, tempData)

	}
	returnedData, _ := json.Marshal(finData)
	return returnedData
}
