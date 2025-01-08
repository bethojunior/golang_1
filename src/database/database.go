package database

import (
	"uno/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func MigrateDatabase() error {
	DB.AutoMigrate(&models.BlogStructModel{})
	DB.AutoMigrate(&models.UserStructModel{})
	return nil
}
