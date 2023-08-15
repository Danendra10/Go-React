package database

import (
	"fmt"
	"log"
	"os"

	"github.com/danendra10/gowlang-first/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = database

	tables := []interface{}{
		&models.User{},
		&models.Post{},
	}

	for _, table := range tables {
		result := database.AutoMigrate(table)
		if result == nil {
			fmt.Printf("Successfully migrated table: %T\n", table)
		}
	}
}
