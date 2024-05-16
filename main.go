package main

// import ./utils/db_utils.go
import (
	"log"
	"product_app/utils"
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

}
