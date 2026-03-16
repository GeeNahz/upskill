package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	// use middleware
	// recover middleware will return a 500 error if a panic occurs
	app.Use(recover.New())
	// middleware for all routes
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Api-Request", "true")
		fmt.Println("Matched all routes.")
		return c.Next()
	})
	// middleware for /api routes
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Matched api routes.")
		return c.Next()
	})

	// get req
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// get req with param
	app.Get("/name::value", func(c *fiber.Ctx) error {
		return c.SendString("Hello, " + c.Params("value"))
	})

	// get req with api with param
	app.Get("/api/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello, " + c.Params("name") + " from /api route")
		} else {
			return c.SendString("Hello, from /api route")
		}
		// return c.SendString("Hello, " + c.Params("value"))
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("This panic is caused by fiber")
	})

	log.Fatal(app.Listen(":3000"))
}
