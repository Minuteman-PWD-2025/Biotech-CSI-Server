package main

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
	// Start the HTTP server on port 8080
	fmt.Printf("Starting Server...\n")
	http.HandleFunc("/", getRoot)
	http.ListenAndServe(":8080", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("Received a POST request")

		// Read the request body
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		data := ""
		for i := 0; i < int(r.ContentLength); i++ {
			data += string(int(reqBody[i]))
		}

		// Check if it's an edit request and if the file already exists
		isEdit := r.FormValue("Edit")
		doesAlreadyExist := AffirmExistanceOfFile(r)
		fmt.Println(doesAlreadyExist)

		if isEdit == "true" {
			if doesAlreadyExist {
				// If editing an existing file, write the new data to it
				os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)
			} else {
				// If the file does not exist, create a new file
				fmt.Println("This file does not exist yet")
				os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)
			}
		} else {
			if doesAlreadyExist {
				// Editing a file that already exists
				os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)
			} else {
				// Creating a new file
				os.WriteFile("./Data/"+r.FormValue("id")+".json", []byte(data), 0644)
			}
		}

		// For removing a file
		if r.FormValue("Deletion") == "true" {
			// Remove the file if requested
			os.Remove("./Data/" + r.FormValue("id") + ".json")
		}

	case "GET":
		if r.FormValue("id") != "" {
			// Serve the JSON file based on the 'id' query parameter
			http.ServeFile(w, r, "./Data/"+string(r.FormValue("id"))+".json")
		}
		if r.FormValue("WhichKey") != "" {
			if r.FormValue("TargetKey") == "|ALL|" {
				// Handle special keyword
			}
			TargetCultures := GetAllRowsWithVal(r)
			os.WriteFile("QuickData", []byte(TargetCultures), 0644)
			http.ServeFile(w, r, "QuickData")
			os.Remove("QuickData")
		}
	}
}

func AffirmExistanceOfFile(r *http.Request) bool {
	// Check if a file exists
	if _, err := os.Stat("./Data/" + r.FormValue("id") + ".json"); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic("Unexpected error occurred")
	}
}

func GetAllRowsWithVal(r *http.Request) string {
	// Check if the "TargetKey" query parameter is provided in the request
	if r.FormValue("TargetKey") != "" {
		alls := "" // Initialize an empty string to store the results

		// Check if the "TargetKey" is not "|ALL|"
		if r.FormValue("TargetKey") != "|ALL|" {
			// Read the list of files in the 'Data' directory
			dir, err := os.ReadDir("./Data")
			if err != nil {
				panic(err)
			}

			// Iterate through the files in the directory
			for _, e := range dir {
				// Read the content of each file
				content, err := os.ReadFile("./Data/" + string(e.Name()))
				if err != nil {
					log.Fatal("Error when opening file: ", err)
				}

				// Parse the JSON content into a map
				var payload map[string]string
				err = json.Unmarshal(content, &payload)
				if err != nil {
					log.Fatal("Error during Unmarshal(): ", err)
				}

				// Check if the value of "WhichKey" matches the "TargetKey"
				if payload[r.FormValue("WhichKey")] == r.FormValue("TargetKey") {
					// Append the "name" value from the JSON data to the result string
					alls += payload["name"] + "\n"
				}
			}
		} else {
			// If "TargetKey" is "|ALL|", retrieve all values of "WhichKey"
			dir, err := os.ReadDir("./Data")
			if err != nil {
				panic(err)
			}

			// Iterate through the files in the directory
			for _, e := range dir {
				// Read the content of each file
				content, err := os.ReadFile("./Data/" + string(e.Name()))
				if err != nil {
					log.Fatal("Error when opening file: ", err)
				}

				// Parse the JSON content into a map
				var payload map[string]string
				err = json.Unmarshal(content, &payload)
				if err != nil {
					log.Fatal("Error during Unmarshal(): ", err)
				}

				// Append the value of "WhichKey" from the JSON data to the result string
				alls += payload[r.FormValue("WhichKey")] + "\n"
			}
		}

		// Check if any matches were found, if not, return a message
		if alls == "" {
			alls = "No Matches Found\n"
		}
		return alls
	}

	// Return an "Invalid request" message if "TargetKey" is not provided in the request
	return "Invalid request"
}
