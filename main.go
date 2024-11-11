package main

import (
	"firstgoserver/database"
	"firstgoserver/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := database.ConnectDB()

	if err != nil {
		fmt.Println("Failed to connect to database: %v", err)
		return
	}

	app := fiber.New()

	router.RouteHandler(app, db)

	app.Listen(":8080")
}
