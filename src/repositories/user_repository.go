package repositories

import (
	"encoding/json"
	"net/http"

	"uno/src/structs"
	// "uno/src/database"

	"gorm.io/gorm"
	
)

var db *gorm.DB

// func InitializeDatabase(dsn string) error {
// 	database.InitDatabase();
// }

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user structs.UserStruct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []structs.UserStruct
	if err := db.Find(&users).Error; err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	var user structs.UserStruct
	if err := db.First(&user, id).Error; err == gorm.ErrRecordNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	var user structs.UserStruct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//user.ID = id // Garantir que o ID seja definido
	if err := db.Model(&user).Updates(user).Error; err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário atualizado"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	if err := db.Delete(&structs.UserStruct{}, id).Error; err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
