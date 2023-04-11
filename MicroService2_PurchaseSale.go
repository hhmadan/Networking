package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type Stock struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func main() {
	http.HandleFunc("/purchase", handlePurchase)
	http.HandleFunc("/sale", handleSale)
	http.ListenAndServe(":8081", nil)
}

func handlePurchase(w http.ResponseWriter, r *http.Request) {
	var stock Stock
	json.NewDecoder(r.Body).Decode(&stock)

	// Make a request to the stock management API to add stock
	reqBody, _ := json.Marshal(stock)
	_, err := http.Post("http://localhost:8080/stocks", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to add stock", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleSale(w http.ResponseWriter, r *http.Request) {
	var stock Stock
	json.NewDecoder(r.Body).Decode(&stock)

	// Make a request to the stock management API to get stock
	res, err := http.Get("http://localhost:8080/stocks/" + strconv.Itoa(stock.ProductID))
	if err != nil {
		http.Error(w, "Failed to get stock", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var currentStock Stock
	json.NewDecoder(res.Body).Decode(&currentStock)

	// Check if there is enough stock
	if currentStock.Quantity < stock.Quantity {
		http.Error(w, "Not enough stock", http.StatusBadRequest)
		return
	}

	// Update the stock and make a request to the stock management API
	currentStock.Quantity -= stock.Quantity
	reqBody, _ := json.Marshal(currentStock)
	req, _ := http.NewRequest("PUT", "http://localhost:8080/stocks/"+strconv.Itoa(stock.ProductID), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		http.Error(w, "Failed to update stock", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
