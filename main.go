package main

import (
	"firstgoserver/models"
	"firstgoserver/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func insertRows(db *gorm.DB, i int) {
	newPost := models.Post{
		Title:       fmt.Sprintf("Title %d", i),
		Description: fmt.Sprintf("Description %d", i),
		UserID:      1,
	}

	if result := db.Create(&newPost); result.Error != nil {
		log.Println("Result Error: ", result.Error)
	}
	fmt.Println("Post Created", newPost.ID)
}

func main() {

	app := fiber.New()

	dsn := "user=root password=root dbname=mytestdb host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router.RouteHandler(app, db)

	// for i := 2376576; i < 20000000; i++ {
	// 	insertRows(db, i)
	// }

	log.Fatal(app.Listen(":8080"))

	// app.Get("/", func(c *fiber.Ctx) error {

	// 	var user User
	// 	result := db.Preload("Posts").First(&user, "username = ?", "mannu")
	// 	// Query by email
	// 	if result.Error != nil {
	// 		log.Println("Error:", result.Error)
	// 	} else {
	// 		log.Println(user)
	// 		// log.Printf("User: ID=%d, Name=%s, Username=%s\n", user.ID, user.Firstname, user.Username)
	// 	}

	// 	var user_data map[string]interface{}
	// 	var Posts_data []map[string]interface{}

	// 	for _, post := range user.Posts {
	// 		d := map[string]interface{}{
	// 			"id":          post.ID,
	// 			"title":       post.Title,
	// 			"description": post.Description,
	// 			// "password":  user.password,
	// 		}
	// 		Posts_data = append(Posts_data, d)
	// 	}
	// 	d := map[string]interface{}{
	// 		"id":        user.ID,
	// 		"firstname": user.Firstname,
	// 		"lastname":  user.Lastname,
	// 		"username":  user.Username,
	// 		"posts":     Posts_data,
	// 		// "pass
	// 	}
	// 	user_data = d

	// 	return c.Status(200).JSON(user_data)
	// })

	// fmt.Println(userData)
}
