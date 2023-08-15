package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danendra10/gowlang-first/database"
	"github.com/danendra10/gowlang-first/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
	fmt.Println("Running on port " + port)
}
