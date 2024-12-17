package controller

import (
	"firstgoserver/models"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatPost(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var requestdata struct {
			Title       string
			Description string
			UserID      uint `json:"user"`
		}
		// Parse the JSON body into the requestdata struct
		err := c.BodyParser(&requestdata)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Create a new Post instance
		newPost := models.Post{
			Title:       requestdata.Title,
			Description: requestdata.Description,
			UserID:      requestdata.UserID,
		}

		// Insert the new post into the database
		if result := db.Create(&newPost); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to create post",
			})
		}

		return c.Status(200).JSON(newPost)
	}
}

func UpdatePost(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		post_id := c.Params("id", "0")

		var requestdata struct {
			Title       string
			Description string
			UserID      uint `json:"user"`
		}

		// Parse the JSON body into the requestdata struct
		err := c.BodyParser(&requestdata)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Find the post with the specified ID
		var post models.Post
		if result := db.First(&post, post_id); result.Error != nil {
			// If the post with the given ID does not exist, return a 404 error
			return c.Status(404).JSON(fiber.Map{
				"error": "Post not found",
			})
		}

		// Update only the fields that are provided in the request
		if requestdata.Title != "" {
			post.Title = requestdata.Title
		}
		if requestdata.Description != "" {
			post.Description = requestdata.Description
		}
		if requestdata.UserID != 0 {
			post.UserID = requestdata.UserID
		}

		// Save the updated post to the database
		if result := db.Save(&post); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to update post",
			})
		}

		return c.Status(200).JSON(post)
	}
}

func GetAllPosts(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		// Get the initial CPU time
		start := time.Now()

		var postsList []map[string]interface{}

		var post []models.Post

		result := db.Where("user_id = ? AND id BETWEEN ? AND ?", 1, 0, 595765).
			Order("id ASC").
			Find(&post)
		if result.Error != nil {
			log.Println("Error:", result.Error)
		}
		for _, post := range post {
			d := map[string]interface{}{
				"id":          post.ID,
				"title":       post.Title,
				"description": post.Description,
				"userId":      post.UserID,
			}
			postsList = append(postsList, d)
		}

		// Calculate the time spent on the task
		duration := time.Since(start)
		fmt.Printf("API processing took: %v\n", duration)

		// Get CPU stats (user and system)
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		fmt.Printf("Alloc = %v MiB", memStats.Alloc/1024/1024)

		return c.Status(200).JSON(postsList)
	}
}

func GetPostsDetails(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		post_id := c.Params("id", "0")

		var posts_data map[string]interface{}

		var post models.Post

		result := db.First(&post, post_id)
		if result.Error != nil {
			log.Println("Error:", result.Error)
			return c.Status(404).JSON(
				map[string]interface{}{
					"err": result.Error.Error(),
				})
		}

		d := map[string]interface{}{
			"id":          post.ID,
			"title":       post.Title,
			"description": post.Description,
		}

		posts_data = d

		return c.Status(200).JSON(posts_data)
	}
}

func GetUserPosts(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		user_id := c.Params("id", "0")

		var user_data map[string]interface{}
		var posts_data []map[string]interface{}

		var user models.User
		result := db.Preload("Posts").First(&user, user_id)
		// Query by email
		if result.Error != nil {
			log.Println("Error:", result.Error)
			return c.Status(404).JSON(
				map[string]interface{}{
					"err": result.Error.Error(),
				})
		}

		for _, post := range user.Posts {
			d := map[string]interface{}{
				"id":          post.ID,
				"title":       post.Title,
				"description": post.Description,
				"user":        post.UserID,
			}
			posts_data = append(posts_data, d)
		}
		d := map[string]interface{}{
			"id":        user.ID,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"username":  user.Username,
			"posts":     posts_data,
		}

		user_data = d

		return c.Status(200).JSON(user_data)
	}
}
