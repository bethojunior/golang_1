package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"uno/src/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterUserRoutes(r)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
