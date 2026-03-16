package main

import (
	"log"
	"todo/src"
	db "todo/src/db"
)

func main() {
	app := src.SetupApp()

	database := db.ConnectDB()
	database.MakeMigrations()

	port := ":3000"
	log.Println("Server Started on Port" + port)
	app.Listen(port)
}
