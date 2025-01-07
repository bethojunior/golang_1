package routes

import (
	"github.com/gorilla/mux"
	"uno/src/handlers"
)

func RegisterUserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", handlers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/", handlers.GetUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.GetUserByID).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")
}
