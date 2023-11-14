package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var users map[string]string
var tokens []string

func main() {
	EnableServer()

	users = make(map[string]string)

	users = map[string]string{
		"email":     "pass",
		"ibroomell": "1234",
		"dledger":   "1234",
	}

	// starts new goroutine for func()
	go func() {
		var lastTime = time.Now()

		for {
			if time.Since(lastTime) >= time.Hour {
				lastTime = time.Now()
				log("invalid tokens removed")
				// check for invalid tokens within SQL database
			}
		}
	}()

	// debug
	go func() {
		for {
			fmt.Scanln()
			fmt.Println(tokens)
		}
	}()

	// Start the HTTP server on port 8080
	log("Starting Server...")
	http.HandleFunc("/api", getRoot)
	http.ListenAndServe(":8080", nil)
}

// called when an application makes a request to server,
// serves relevant files and makes relevant changes to data
func getRoot(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		log("recieved post request")

		// if email and password provided in url query
		if r.FormValue("email") != "" && r.FormValue("password") != "" {
			// initialize email and password as query values
			email := r.FormValue("email")
			password := r.FormValue("password")

			// initialize error for error handling
			var err error

			// validate login information, if valid update token array with new token
			tokens, err = ValidateLogin(users, tokens, email, password)
			if err != nil {
				log("error logging in: " + err.Error())
				return
			}

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

					AddNew(table, finStringC, finStringV)
				}
			}
		}

	case "GET":
		log("recieved get request")
		if r.FormValue("token") != "" {
			if r.FormValue("table") != "" {

				rows := GetTable(r.FormValue("table"))
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
