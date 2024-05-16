package main

// import ./utils/db_utils.go
import (
	"log"
	"net/http"

	"product_app/handlers"
	"product_app/utils"

	"github.com/gorilla/mux"
)

var db = utils.NewDatabase()

func main() {
	//establish connection to db
	connStr := "user=anhng password=password dbname=product_app sslmode=disable"
	err := db.Init(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Run db migrations
	// err = db.RunMigrations("file://db/migrations")

	// if err != nil {
	// 	log.Fatalf("Failed to run migrations: %v", err)
	// }

	// Insert mock data
	// err = utils.InsertMockData(db.DB)
	// if err != nil {
	// 	log.Fatalf("Failed to insert mock data: %v", err)
	// }

	// Delete mock data
	// err = utils.DeleteMockData(db.DB)
	// if err != nil {
	// 	log.Fatalf("Failed to delete mock data: %v", err)
	// }

	router := mux.NewRouter()
	productHandler := &handlers.ProductHandler{DB: db.DB}
	categoryHandler := &handlers.CategoryHandler{DB: db.DB}

	router.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/categories", categoryHandler.GetCategories).Methods("GET")
	router.HandleFunc("/", productHandler.ListProducts).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
