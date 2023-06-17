package config

import (
	"log"
	"os"
	"posts/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	if err := DB.AutoMigrate(&models.User{}, &models.Client{}); err != nil {
		err.Error()
	}
}
