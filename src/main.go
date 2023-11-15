package main

import (
	"fmt"
	"os"

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
		"email":      "pass",
		"ibroomell":  "1234",
		"drewledger": "1624",
	}

	// starts new goroutine for func()
	go func() {
		var lastTime = time.Now()

		for {
			if time.Since(lastTime) >= time.Hour {
				lastTime = time.Now()
				log(true, "invalid tokens removed")
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
	log(false, "Starting Server...")
	http.HandleFunc("/api", getRoot)
	http.ListenAndServe(":8080", nil)
}

// Helper function for GET request modularity
func handleLoginRequest(r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	var err error

	fmt.Printf("tokens before login: %v\n", tokens)
	fmt.Println(len(users), "Users are connected")
	fmt.Printf("-------------------------------------\n")
	for userEmail, userPassword := range users {
		fmt.Printf("User: %s - Password: %s\n", userEmail, userPassword)
	}
	fmt.Printf("-------------------------------------\n")

	tokens, err = ValidateLogin(users, tokens, email, password)
	if err != nil {
		log(true, "error logging in: %s\n", err.Error())
		return
	}

	fmt.Printf("tokens after login: %s\n", strings.Join(tokens, " "))
}

// called when an application makes a request to server,
// serves relevant files and makes relevant changes to data
func getRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		log(false, "recieved post request")

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
				log(true, "error logging in: "+err.Error())
				return
			}

		}
		if r.FormValue("token") != "" {
			if r.FormValue("table") != "" && r.FormValue("insert") != "" {
				table := r.FormValue("table")
				dat := (r.FormValue("insert"))
				splitDat := strings.Split(dat, sep)
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

	case "GET":
		log(true, "recieved get request")
		if r.FormValue("token") != "" {
			if r.FormValue("table") != "" {
				returnedData := FormatTableToJSON(r.FormValue("table"))
				os.WriteFile("data.json", returnedData, 0644)
				http.ServeFile(w, r, "data.json")
			}
		} else {
			if r.FormValue("email") != "" && r.FormValue("password") != "" {
				handleLoginRequest(r)
			}
		}
	case "PUT":
		update := r.FormValue("update")
		indivUpdate := strings.Split(update, sep)
		where := r.FormValue("where")
		indivWhere := strings.Split(where, sep)
		err := AlterThing(r.FormValue("table"), indivUpdate, indivWhere)
		if err != nil {
			log(true, "error in PUT request: "+err.Error())
		}

	//Authentication
	case "DELETE":
		table := r.FormValue("table")
		del := r.FormValue("where")
		indivDells := strings.Split(del, sep)
		err := DeleteRow(table, indivDells)
		if err != nil {
			log(true, "error in DELETE request: "+err.Error())
		}
	}
}
