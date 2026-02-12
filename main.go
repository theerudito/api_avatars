package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()

	router(app)

	port := os.Getenv("PortServer")

	if port == "" {
		port = "1001"
	}

	log.Fatal(app.Listen(":" + port))

}
