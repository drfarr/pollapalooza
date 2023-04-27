package database

import (
	"log"
	"os"
	"pollapalooza/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Connect() {

	var err error
	dsn := goDotEnvVariable("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_ = DB
	if err != nil {
		log.Fatalln(err)
	}

}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
