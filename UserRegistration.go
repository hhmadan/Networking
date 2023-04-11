package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

var users []User

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/userregistration", handlePostRequest)
	http.HandleFunc("/userinfo", userInfo)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func handleRoot(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("---------- WELCOME TO USER REGISTRATION ----------"))
}

func handlePostRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var user User
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "Invalid request: %v", err)
			return
		}
		users = append(users, user)
		writer.WriteHeader(http.StatusCreated)
		fmt.Println("Item was successfully added to slice.")
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Invalid Request-Method Not Found")
	}
}

func userInfo(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		json.NewEncoder(writer).Encode(users)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(writer, "Invalid method")
	}
}
