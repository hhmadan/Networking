package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Fruits struct {
	Name, Color string
	Id, Price   int
}

var fruits []Fruits

func main() {

	http.HandleFunc("/", handleDisplayPage)
	http.HandleFunc("/data", func(w http.ResponseWriter, rq *http.Request) {
		if rq.Method == http.MethodPost {
			json.NewDecoder(rq.Body).Decode(&fruits)
		}
	})

	http.ListenAndServe(":8000", nil)
}
func handleDisplayPage(w http.ResponseWriter, rq *http.Request) {
	//w.Write([]byte("This is Fruits Data Web Page"))
	fmt.Fprint(w, "This is Fruits Data Web Page")
}

func deleteDetailsMethod(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		id := request.URL.Query().Get("id")
		if id == "" {
			http.Error(writer, "Invalid Request", http.StatusBadRequest)
			return
		}

		idRcvd, _ := strconv.Atoi(id)

		for i, r := range fruits {
			if r.Id == idRcvd {
				fruits = append(fruits[:i], fruits[i+1:]...)

			}
		}
		writer.WriteHeader(http.StatusOK)
	}
}
