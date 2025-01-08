package controllers

import (
	"net/http"

	"uno/src/handlers"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	handlers.CreateUser(w, r)
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	handlers.GetUsers(w, r)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	handlers.GetUserByID(w, r)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateUser(w, r)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteUser(w, r)
}
