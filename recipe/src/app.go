package src

import (
	// "log"

	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Recipe App",
	})

	return app
}
