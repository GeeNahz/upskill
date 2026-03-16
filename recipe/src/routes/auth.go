package routes

import (
	"github.com/geenahz/recipe/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	// authenticated
	app.Post("/register", handlers.RegisterHandler)
	// unauthenticated
	app.Post("/login", handlers.LoginHandler)
}
