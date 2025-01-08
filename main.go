package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"uno/src/database"
	"uno/src/routes"
)

func main() {
	r := mux.NewRouter()

	dsn := "host=localhost user=root password=password dbname=db port=5432 sslmode=disable"

	if err := database.InitializeDatabase(dsn); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// play migrations
	if err := database.MigrateDatabase(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	routes.RegisterUserRoutes(r)

	log.Println("Server listening on port 8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", r))
}
