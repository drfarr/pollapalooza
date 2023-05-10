package database

import (
	"fmt"
	"log"
	"os"
	"pollapalooza/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var err error
	dsn := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.Poll{}, &models.User{})
	fmt.Println("Database Migrated")

}
