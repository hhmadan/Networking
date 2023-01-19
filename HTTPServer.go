package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleHome)

	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second) // 5 seconds delay
	fmt.Fprint(w, "Hello World!")
}
