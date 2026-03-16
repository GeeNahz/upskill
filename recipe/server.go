package main

import (
	"fmt"
	"os"

	"github.com/geenahz/recipe/src"
	db "github.com/geenahz/recipe/src/database"
	"github.com/geenahz/recipe/src/routes"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	app := src.SetupApp()
	v1 := app.Group("/api/v1")

	database := db.ConnectDB()
	database.Migrate()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	routes.AuthRoutes(v1.Group("/auth"))

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
