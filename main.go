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
		isEdit := r.FormValue("Edit")
		doesAlreadyExist := AffirmExistanceOfFile(r)
		fmt.Println(doesAlreadyExist)
		if isEdit == "true" {
			if doesAlreadyExist { //Means you are editing a file that exists so everything is ok
				//On the frontend we will need to grab the data of the file, and then fill it in, such that they only need to edit things

				os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)

			} else { //Uh oh, the file does not exist, will ask if you want to create it
				fmt.Println("This file does not exist yet")
				os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)
			}
		} else {
			if doesAlreadyExist { //uh oh, you are trying to creat a new file that already exists
				//Replace with some proc for confirmation
				fmt.Println("This file already exists, would you like to edit it instead?")
				os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)

			} else { //You are creating a new file that does not exist yet, so go ahead with it

				os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)
			}
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
