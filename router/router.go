package router

import (
	"firstgoserver/controller"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func RouteHandler(app *fiber.App, db *gorm.DB) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Wooo Hooo!, You are connected now")
	})

	// User APIS
	app.Get("/api/v2/users/:id/", controller.GetUserDetails(db))
	app.Get("/api/v2/users/", controller.GetAllUsers(db))
	// Psts APIS
	app.Get("/api/v2/posts/", controller.GetAllPosts(db))
	app.Post("/api/v2/posts/", controller.CreatPost(db))
	app.Patch("/api/v2/posts/:id/", controller.UpdatePost(db))
	app.Get("/api/v2/posts/:id", controller.GetPostsDetails(db))
	app.Get("/api/v2/user-posts/:id", controller.GetUserPosts(db))

}
