package handlers

import (
	// "fmt"
	"http/library/src/helpers"
	"http/library/src/models"
	"http/library/src/repo"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler interface {
	registerUser(c *fiber.Ctx) error
	loginUser(c *fiber.Ctx) error
}

func Authhandlers(route fiber.Router, db *gorm.DB) {
	route.Post("/register", func(c *fiber.Ctx) error {
		return registerUser(c, db)
	})

	route.Post("/login", func(c *fiber.Ctx) error {
		return loginUser(c, db)
	})
}

func registerUser(c *fiber.Ctx, db *gorm.DB) error {
	user := &models.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse body", "message": err.Error(),
		})
	}
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error", "message": "Username and password are required",
		})
	}

	hashed, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Hash error", "message": err.Error(),
		})
	}

	user.Password = hashed
	repo.CreateUser(user, db)

	token, err := helpers.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token error", "message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: !c.IsFromLocal(),
		Secure:   !c.IsFromLocal(),
		MaxAge:   3600 * 24 * 7,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func loginUser(c *fiber.Ctx, db *gorm.DB) error {
	dbUser := new(models.User)
	authUser := &models.User{}
	if err := c.BodyParser(authUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse body", "message": err.Error(),
		})
	}
	if authUser.Username == "" || authUser.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error", "message": "Username and password are required",
		})
	}

	// get user
	repo.GetUserByUsername(authUser.Username, db, dbUser)
	if dbUser.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Not found error", "message": "Invalid username or password",
		})
	}

	if err := helpers.CompareHashAndPassword(dbUser.Password, authUser.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized error", "message": "Invalid username or password",
		})
	}

	token, err := helpers.GenerateToken(authUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token error", "message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: !c.IsFromLocal(),
		Secure:   !c.IsFromLocal(),
		MaxAge:   3600 * 24 * 7,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}
