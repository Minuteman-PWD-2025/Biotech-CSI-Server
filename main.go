package main

//this is in need of error handling in some places
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

		//Optimize this to check for deletion before doing the writing
		//Actually, by writing before deleting, it automatically handles the possible issue of the file to delete not existing, so that is nice
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
		//For removing from the dat
		if r.FormValue("Deletion") == "Yes" {
			os.Remove("./Data/" + r.FormValue("id") + ".json")
		}
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
