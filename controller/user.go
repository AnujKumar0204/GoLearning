package controller

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	id        int
	firstname string
	lastname  string
	username  string
	password  string
}

func GetUserDetails(db *sql.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {

		user_id := c.Params("id", "0")

		var data []map[string]interface{}

		var userList []User
		rows, err := db.Query("select id,firstname,lastname,username from account_useraccount where id = $1", user_id)

		if err != nil {
			fmt.Println(err)
			return c.Status(400).SendString("err")
		}
		defer rows.Close()

		// for rows.Next() {

		// 	var id int
		// 	var firstname, lastname, username, password string

		// 	errrr := rows.Scan(&id, &firstname, &lastname, &username, &password)
		// 	if errrr != nil {
		// 		log.Fatal(errrr)
		// 	}

		// 	fmt.Printf("ID: %d, Name: %s, Email: %s, Username: %s, Password: %s\n", id, firstname, lastname, username, password)
		// }

		for rows.Next() {
			var user User
			err := rows.Scan(
				&user.id,
				&user.firstname,
				&user.lastname,
				&user.username,
				// &user.password,
			)
			if err != nil {
				log.Fatal(err)
			}
			// d := map[string]interface{}{
			// 	"id":        user.id,
			// 	"firstname": user.firstname,
			// 	"lastname":  user.lastname,
			// 	"username":  user.username,
			// 	"password":  user.password,
			// }
			userList = append(userList, user)
		}

		for _, user := range userList {
			d := map[string]interface{}{
				"id":        user.id,
				"firstname": user.firstname,
				"lastname":  user.lastname,
				"username":  user.username,
				// "password":  user.password,
			}
			data = append(data, d)
		}

		return c.Status(200).JSON(data)
	}
}
