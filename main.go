package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Printf("Hello Worlds")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/Hello", GetHello)
	http.ListenAndServe(":1234", nil)

}
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World")
	io.WriteString(w, "Hello\n")

}
