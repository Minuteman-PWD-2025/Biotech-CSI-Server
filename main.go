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

			//This line isnt needed and will be removed in my next push
			//io.WriteString(w, payload.Name)
			//f, err := os.Open()
			http.ServeFile(w, r, "./Data/"+string(r.FormValue("id"))+".json")
		}
		if r.FormValue("WhichKey") != "" {
			//Any special keywords will be in all caps and enclosed in pipes
			if r.FormValue("TargetKey") == "|ALL|" {

			}
			TargetCultures := GetAllRowsWithVal(r)
			os.WriteFile("QuickData", []byte(TargetCultures), 0644)
			http.ServeFile(w, r, "QuickData")
			os.Remove("QuickData")
		}
		//Get all the items with a matching

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
func GetAllRowsWithVal(r *http.Request) string { //Just return it all for a string right now
	if r.FormValue("TargetKey") != "" {
		alls := ""
		if r.FormValue("TargetKey") != "|ALL|" {
			dir, err := os.ReadDir("./Data")
			if err != nil {
				panic(err)
			}

			for _, e := range dir {

				content, err := os.ReadFile("./Data/" + string(e.Name()))

				if err != nil {
					log.Fatal("Error when opening file: ", err)
				}

				var payload map[string]string
				err = json.Unmarshal(content, &payload)
				if err != nil {
					log.Fatal("Error during Unmarshal(): ", err)
				}
				if payload[r.FormValue("WhichKey")] == r.FormValue("TargetKey") {
					alls += payload["name"] + "\n"
				}

			}
		} else {
			dir, err := os.ReadDir("./Data")
			if err != nil {
				panic(err)
			}

			for _, e := range dir {

				content, err := os.ReadFile("./Data/" + string(e.Name()))

				if err != nil {
					log.Fatal("Error when opening file: ", err)
				}

				var payload map[string]string
				err = json.Unmarshal(content, &payload)
				if err != nil {
					log.Fatal("Error during Unmarshal(): ", err)
				}

				alls += payload[r.FormValue("WhichKey")] + "\n"

			}
		}
		if alls == "" {
			alls = "No Matches Found\n"
		}
		return alls

	}
	return "How did you get here"
}
