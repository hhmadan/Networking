package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Struct creation for stock
type Stock struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// slice of stock
var stocks []Stock

// main method addition
func main() {

	// Ready made data
	stocks = []Stock{
		{ProductID: 1, Quantity: 10},
		{ProductID: 2, Quantity: 20},
	}

	http.HandleFunc("/stocks", handleStocks) //endpoint for multiple stock objects
	http.HandleFunc("/stocks/", handleStock) //endpoint for particular stock object
	http.ListenAndServe(":8080", nil)
}

/*
	handleStocks func
*/

func handleStocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle stocks multi item api was called")
	switch r.Method {
	case "GET":
		handleGetStocks(w, r)
	case "POST":
		handlePostStock(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// handle get stocks will return the all stocks

func handleGetStocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stocks)

}

// handle Post Stock
func handlePostStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle stock single item api was called")
	var stock Stock
	json.NewDecoder(r.Body).Decode(&stock)
	stocks = append(stocks, stock)
	json.NewEncoder(w).Encode(stock)

}

// another endpoint 2 method
// it will handle all the
// request for particular object

func handleStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/stocks/"):])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i, stock := range stocks {
		if stock.ProductID == id {
			switch r.Method {
			case "GET":
				json.NewEncoder(w).Encode(stock)
			case "PUT":
				var newStock Stock
				json.NewDecoder(r.Body).Decode(&newStock)
				newStock.ProductID = id
				stocks[i] = newStock
				json.NewEncoder(w).Encode(newStock)
			case "DELETE":
				stocks = append(stocks[:i], stocks[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
			default:
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			}
			return
		}
	}

	http.Error(w, "Stock not found", http.StatusNotFound)
}
