package main

//this is in need of error handling in some places
import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	//Currently can only get the server working on 8080
	fmt.Printf("Starting Server...\n")
	http.HandleFunc("/", getRoot)

	http.ListenAndServe(":8080", nil)

}
func getRoot(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		//by writing before deleting, it automatically handles the possible issue of the file to delete not existing, so that is nice
		fmt.Println("Got A post")

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		data := ""
		for i := 0; i < int(r.ContentLength); i++ {
			data += string(int(reqBody[i]))
		}

		//Creates or updates the file
		os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)

		doesAlreadyExist := AffirmExistanceOfFile(r)
		fmt.Println(doesAlreadyExist)

		if doesAlreadyExist {
			//Need something to proc for confirmation here
			fmt.Println("Already Exists!")
			os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)

		} else {
			os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)
		}

		//For removing from the dat
		if r.FormValue("Deletion") == "true" {
			os.Remove("./Data/" + r.FormValue("id") + ".json")
		}
	case "GET":
		if r.FormValue("id") != "" {
			//Read Json from the file requested, based on id
			content, err := os.ReadFile("./Data/" + string(r.FormValue("id")) + ".json")
			fmt.Printf(r.FormValue("id") + "\n")
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}
			//Convert it to stringstuff
			var payload Data
			err = json.Unmarshal(content, &payload)
			if err != nil {
				log.Fatal("Error during Unmarshal(): ", err)
			}
			//This line isnt needed and will be removed in my next push
			//io.WriteString(w, payload.Name)
		}

	}
}
func AffirmExistanceOfFile(r *http.Request) bool {
	if _, err := os.Stat("./Data/" + r.FormValue("id") + ".json"); err == nil {
		return true

	} else if errors.Is(err, os.ErrNotExist) {
		return false

	} else {
		panic("What are you doing here how did you Schrodinger")

	}
}
