package handlers

import (
	"fmt"

	"github.com/geenahz/recipe/src/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create user: %v", err.Error()),
		})
	}

	// validate user

	// create record in database
	// return user and status code

	return c.SendString("Register route")
}

func LoginHandler(c *fiber.Ctx) error {
	return c.SendString("Login route")
}
