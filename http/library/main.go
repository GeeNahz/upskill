package main

import (
	"http/library/src/database"
	"http/library/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialise the database
	db := database.InitializeDB()

	// Create the Fiber App Instance
	app := fiber.New(fiber.Config{
		AppName: "Library API",
	})
	api := app.Group("/api")

	apiV1 := api.Group("/v1")

	// Define the Auth routes. Those will be public
	handlers.Authhandlers(apiV1.Group("/auth"), db)

	// Start server on port 4000
	app.Listen(":4000")
}
