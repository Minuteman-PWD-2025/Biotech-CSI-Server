package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Printf("Starting Server...\n")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/Hello", GetHello)
	http.ListenAndServe(":1234", nil)

}
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	switch r.Method {
	case "POST":

	case "GET":

	}
}
func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World")
	io.WriteString(w, "Hello\n")

}
