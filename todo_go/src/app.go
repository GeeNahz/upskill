package src

import (
	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Todo API made with Go")
	})

	return app
}
