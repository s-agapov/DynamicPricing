package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/duckdb/duckdb-go/v2"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	http.HandleFunc("/products", getProductsHandler)
	fmt.Println("Server started on :8086")
	http.ListenAndServe(":8086", nil)
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Connect to DuckDB database
	conn, err := duckdb.Connect("localhost:3142", "products.duckdb") // Replace with your DuckDB connection string
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Execute query to fetch product data
	rows, err := conn.Query("SELECT id, name, price FROM products")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process and marshal product data
	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	// Send JSON response
	json.NewEncoder(w).Encode(products)
}
