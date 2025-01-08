package routes

import (
	"log"
	"uno/src/handlers"

	"github.com/gorilla/mux"
)
									
func RegisterUserRoutes(r *mux.Router) {
	log.Println("Server User Routes 🚀")
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", handlers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/", handlers.GetUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.GetUserByID).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")
}
