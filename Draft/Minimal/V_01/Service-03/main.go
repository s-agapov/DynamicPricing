package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// Product struct representing the data to be stored in the DB
type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var db *gorm.DB

func init() {
	// Connect to SQLite database (creates if it doesn't exist)
	var err error
	db, err = gorm.Open("sqlite3", "products.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the schema (create table if it doesn't exist)
	db.AutoMigrate(&Product{})
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Mock data - replace with actual API call to fetch product data
	products := []Product{
		{Name: "Laptop", Price: 1200},
		{Name: "Smartphone", Price: 800},
		{Name: "Tablet", Price: 350},
	}

	// Process and save products in the database
	for _, p := range products {
		db.Create(&p) // Create a new product record
	}

	fmt.Fprintf(w, "Products saved successfully!")
}

func main() {
	http.HandleFunc("/products", getProductsHandler)
	fmt.Println("Server started on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
