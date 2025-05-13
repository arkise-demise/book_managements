package main

import (
	"book_managements/pkg/config"
	"book_managements/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Connect to the database
	db, err := config.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Create a new router instance
	r := mux.NewRouter()

	// Register routes
	routes.RegisterBookStoreRoutes(r)

	// Handle routes
	http.Handle("/", r)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
