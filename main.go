package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"net/http"
)

var users map[string]string

func main() {
	dosmt()

	// Start the HTTP server on port 8080
	fmt.Printf("Starting Server...\n")
	http.HandleFunc("/", getRoot)
	http.ListenAndServe(":8080", nil)

	users = make(map[string]string)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("recieved post request")

	case "GET":
		fmt.Println("recieved get request")
	}
}

// Add password encryption
func ValidateLogin(email string, pass string) (string, error) {
	_, exists := users[email]
	if !exists {
		users[email] = pass
	}
	return "", nil
}

func GenerateToken(length int) (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length], nil
}
