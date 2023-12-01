package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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

// EnableServer connects to the database and sets the global db variable.
func EnableServer() {
	mydb, err := ConnectToTable()
	if err != nil {
		log(true, err.Error())
		//os.Exit(404)
		return
	}
	db = mydb
}

// ConnectToTable establishes a connection to the database and returns the *sql.DB object.
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

// GetTable retrieves all rows from the specified table.
// If a where clause is provided, it filters the rows based on the condition.
func GetTable(table string, where string) sql.Rows {
	rows, err := db.Query("SELECT * FROM " + RemSpaces(table) + ";")
	if where != "" {
		indivwhere := strings.Split(where, ",")
		rows, err = db.Query("SELECT * FROM " + RemSpaces(table) + " WHERE " + indivwhere[0] + indivwhere[1])
	}
	if err != nil {
		panic(err)
	}
	return *rows
}

// RemSpaces removes spaces from the input query and returns the first part of the string.
// This is done to prevent SQL injections.
func RemSpaces(query string) string {
	splitString := strings.Split(query, " ")
	return splitString[0]
}

// AddNew inserts a new row into the specified table with the provided data.
// It returns the inserted rows and any error encountered.
func AddNew(table string, cols string, Data string) (sql.Rows, error) {
	rows, err := db.Query("INSERT INTO " + table + " " + cols + "\nVALUES " + Data)
	if err != nil {
		err = errors.New("error inserting data: " + err.Error())
	}
	return *rows, err
}

// AlterThing updates the specified table with the given updates and conditions.
// It iterates over the updates and conditions and performs the corresponding SQL update statements.
// Returns an error if any update fails.
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

// DeleteRow deletes rows from the specified table based on the given conditions.
// It iterates over the conditions and performs the corresponding SQL delete statements.
// Returns an error if any delete operation fails.
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

// FormatTableToJSON retrieves rows from the specified table and converts them to JSON format.
// If a where clause is provided it filters the rows based on the condition.
// Returns the JSON data as a byte slice.
func FormatTableToJSON(table string, where string) []byte {
	rows := GetTable(table, where)
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
