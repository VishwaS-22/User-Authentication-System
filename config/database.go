package config

import (
	"Users-CRUD-API/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(GetEnv("DB_Name", "users.db")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}
	log.Println("Database connected successfully.")

	DB.AutoMigrate(&models.User{})
}
