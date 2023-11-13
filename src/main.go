package main

import (
	"fmt"
	"net/http"
	"strings"
)

var users map[string]string
var tokens []string

func main() {
	testConnection()

	users = make(map[string]string)

	users = map[string]string{
		"email":     "pass",
		"ibroomell": "1234",
		"dledger":   "1234",
	}

	// Start the HTTP server on port 8080
	fmt.Printf("Starting Server...\n")
	http.HandleFunc("/api", getRoot)
	http.ListenAndServe(":8080", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	err, db := testConnection()
	switch r.Method {
	case "POST":
		fmt.Println("recieved post request")
		if r.FormValue("token") != "" {
			if r.FormValue("table") != "" && r.FormValue("insert") != "" {
				table := r.FormValue("table")
				dat := (r.FormValue("insert"))
				splitDat := strings.Split(dat, "|")
				finStringC := "("
				finStringV := "("

				for i := 0; i < len(splitDat); i++ {
					if i < len(splitDat)-1 {
						finStringC += strings.Split(splitDat[i], ",")[0] + ", "
						finStringV += strings.Split(splitDat[i], ",")[1] + ", "
					} else {
						finStringC += strings.Split(splitDat[i], ",")[0] + ")"
						finStringV += strings.Split(splitDat[i], ",")[1] + ")"
					}

				}

				AddNew(table, db, finStringC, finStringV)
			}
		}

	case "GET":
		fmt.Println("recieved get request")
		if r.FormValue("token") != "" {
			if r.FormValue("table") != "" {

				if err != nil {
					panic(err)
				}
				rows := GetTable(r.FormValue("table"), db)
				//#region Placeholder
				var user_id string
				var name string

				for rows.Next() {
					rows.Scan(&user_id, &name)
					fmt.Printf("ID: %s\nName: %s\n\n", user_id, name)
				}

				//#endregion
			}
		} else {
			if r.FormValue("email") != "" && r.FormValue("password") != "" {
				email := r.FormValue("email")
				password := r.FormValue("password")

				var err error

				fmt.Print("tokens before login: ")
				fmt.Println(tokens)

				fmt.Print("users: ")
				fmt.Println(users)

				tokens, err = ValidateLogin(users, tokens, email, password)
				if err != nil {
					fmt.Println("error logging in: " + err.Error())
					return
				}

				fmt.Print("tokens after login: ")
				for i := 0; i < len(tokens); i++ {
					fmt.Print(tokens[i] + " ")
				}
				fmt.Println("")
			}
		}

		//Authentication

	}

}
