package config

import (
	"fmt"
	"go-todo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB: ", err)
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate DB: ", err)
	}

	DB = db

	fmt.Println("Connected to DB")
}
