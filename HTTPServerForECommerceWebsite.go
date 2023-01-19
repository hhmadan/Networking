// Write a Go program that creates a HTTP server for a e-commerce website,
// that listens on
// a specific port, and responds to incoming HTTP requests
// by fetching product information from a database and returning it to the
// client in a JSON format.

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("mysql", "root:Hemangi9@root@tcp(127.0.0.1:3306)/eCommerceProductsData")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	http.HandleFunc("/requestedresource", func(w http.ResponseWriter, r *http.Request) {
		// Fetch product from database
		product, _ := readDataFromDB()

		// Convert product to JSON
		productJSON, _ := json.Marshal(product)

		// Write JSON to response
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)
	})

	http.ListenAndServe(":8080", nil)
}

func readDataFromDB() ([]Product, error) {
	var prod []Product

	rows, err := db.Query("SELECT * FROM products;")
	if err != nil {
		return nil, fmt.Errorf("error in query all products: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign record to slice
	for rows.Next() {
		var products Product
		if err := rows.Scan(&products.Id, &products.Name, &products.Description, &products.Price); err != nil {
			return nil, fmt.Errorf("error in query all products: %v", err)
		}
		prod = append(prod, products)
	}
	return prod, nil
}

// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type Product struct {
// 	Id          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	Price       int    `json:"price"`
// }

// var db *sql.DB

// func main() {
// 	var err error
// 	//db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ecommerce")
// 	db, err = sql.Open("mysql", "root:Hemangi9@root@tcp(127.0.0.1:3306)/eCommerceProductsData")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	http.HandleFunc("/products", productsHandler)

// 	fmt.Println("Starting server on port 8080...")
// 	http.ListenAndServe(":8080", nil)
// }

// func productsHandler(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("SELECT * FROM products")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var products []Product
// 	for rows.Next() {
// 		var p Product
// 		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		products = append(products, p)
// 	}

// 	json.NewEncoder(w).Encode(products)
// }
