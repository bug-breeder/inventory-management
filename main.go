package main

import (
	"log"
	"net/http"
	"product_app/handlers"
	"product_app/utils"

	"github.com/gorilla/mux"
)

func main() {
	db := utils.NewDatabase()

	connStr := "user=anhng password=password dbname=product_app sslmode=disable"
	err := db.Init(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Run migrations
	err = db.RunMigrations("file://db/migrations")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	router := mux.NewRouter()
	// Log all requests
	// router.Use(utils.LoggingMiddleware)

	productHandler := &handlers.ProductHandler{DB: db.DB}
	categoryHandler := &handlers.CategoryHandler{DB: db.DB}

	router.HandleFunc("/", productHandler.ListProducts).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/categories", categoryHandler.GetCategories).Methods("GET")
	router.HandleFunc("/add-product", productHandler.ShowAddProductPage).Methods("GET")
	router.HandleFunc("/add-product", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/edit-product/{id}", productHandler.ShowEditProductPage).Methods("GET")
	router.HandleFunc("/edit-product/{id}", productHandler.EditProduct).Methods("PATCH")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
