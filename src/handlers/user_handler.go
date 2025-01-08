package handlers

import (
	"encoding/json"
	"net/http"

	"uno/src/structs"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado"})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []structs.UserStruct{
		{ID: 1, Name: "João", Email: "joao@email.com"},
		{ID: 2, Name: "Maria", Email: "maria@email.com"},
	}
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(structs.UserStruct{ID: 1, Name: "João", Email: "joao@email.com"})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário atualizado"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
