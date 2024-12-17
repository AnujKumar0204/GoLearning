package controller

import (
	"firstgoserver/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {

		var usersList []map[string]interface{}

		var user []models.User

		result := db.Find(&user)
		if result.Error != nil {
			log.Println("Error:", result.Error)
		}

		for _, user := range user {
			d := map[string]interface{}{
				"id":        user.ID,
				"firstname": user.Firstname,
				"lastname":  user.Lastname,
				"username":  user.Username,
			}
			usersList = append(usersList, d)
		}

		return c.Status(200).JSON(usersList)
	}
}

func GetUserDetails(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		user_id := c.Params("id", "0")

		var user_data map[string]interface{}
		// var Posts_data []map[string]interface{}

		var user models.User
		result := db.First(&user, user_id)
		if result.Error != nil {
			log.Println("Error:", result.Error)
			return c.Status(404).JSON(
				map[string]interface{}{
					"err": result.Error.Error(),
				})
		}

		// for _, post := range user.Posts {
		// 	d := map[string]interface{}{
		// 		"id":          post.ID,
		// 		"title":       post.Title,
		// 		"description": post.Description,
		// 		// "password":  user.password,
		// 	}
		// 	Posts_data = append(Posts_data, d)
		// }
		d := map[string]interface{}{
			"id":        user.ID,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"username":  user.Username,
		}
		user_data = d

		return c.Status(200).JSON(user_data)
	}
}
