package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing token",
		})
	}

	tokenStr := strings.Replace(header, "Bearer ", "", 1)
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing token",
		})
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", uint(claims["user_id"].(float64)))

	return c.Next()
}
