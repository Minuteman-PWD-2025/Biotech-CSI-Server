package main

import (
	"fmt"
	"net/http"
	"time"
)

var users map[string]string
var tokens []string

func main() {
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

	case "GET":
		log("recieved get request")

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
		}
	}
}
