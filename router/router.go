package router

import (
	"database/sql"
	"firstgoserver/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteHandler(app *fiber.App, db *sql.DB) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Wooo Hooo!, You are connected now")
	})

	app.Get("/api/v2/get-user/:id/", controller.GetUserDetails(db))

}
