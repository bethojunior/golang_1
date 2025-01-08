package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	log.Println("Banco de dados conectado")
}
