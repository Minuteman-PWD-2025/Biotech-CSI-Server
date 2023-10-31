package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Name         string
	IsCheckedOut bool
	CheckedOutBy string
	TimeOut      string
	Note         string
}

func main() {

	fmt.Printf("Starting Server...\n")
	http.HandleFunc("/", getRoot)

	http.ListenAndServe(":8080", nil)

}
func getRoot(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		fmt.Println("Got A post")

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		data := ""
		for i := 0; i < int(r.ContentLength); i++ {
			data += string(int(reqBody[i]))
		}

		os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)
		//Check if the file already exists, and if it does, "modify"(1) it, else, create and write to it

	case "GET":
		if r.FormValue("id") != "" {
			//Read Json from the file requested
			content, err := ioutil.ReadFile("./Data/" + string(r.FormValue("id")) + ".json")
			fmt.Printf(r.FormValue("id") + "\n")
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}
			var payload Data
			err = json.Unmarshal(content, &payload)
			if err != nil {
				log.Fatal("Error during Unmarshal(): ", err)
			}
			io.WriteString(w, payload.Name)
		}

	}
}

//(1) Actually just rewrite the file with new data, and keeping old data. Query kit for more info.
