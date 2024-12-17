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

	// server := socketio.NewServer(nil)

	// // server.OnConnect("/", func(s socketio.Conn) error {
	// // 	s.SetContext("")
	// // 	fmt.Println("connected:", s.ID())
	// // 	return nil
	// // })

	// // Handle a "msg" event and broadcast a reply to the sender
	// // server.OnEvent("/chat", "joinRoom", func(s socketio.Conn, roomId string) {

	// // 	// The client sent a roomId, we can join the user to that room
	// // 	s.Join(roomId) // Join the requested room
	// // 	s.Emit("reply", "have have connected to room "+roomId)
	// // 	fmt.Println("User wants to join room:", roomId)
	// // })

	// // Handle user connections
	// server.OnConnect("/chat", func(s socketio.Conn) error {
	// 	// fmt.Println("A user connected:", s.ID())

	// 	queryString := s.URL().RawQuery

	// 	// Parse the query string
	// 	values, err := url.ParseQuery(queryString)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// Get the 'roomId' value
	// 	roomId := values.Get("roomId")
	// 	user_name := values.Get("user")
	// 	// fmt.Println("User is:", user)

	// 	// set user for the connection
	// 	user := User{
	// 		roomId: roomId,
	// 		name:   user_name,
	// 	}
	// 	// Associate the user struct with the connection ID
	// 	users.Store(s.ID(), user)

	// 	// Automatically add the user to room 2 upon connection
	// 	if roomId != "" {
	// 		s.Join(roomId) // The user is automatically added to room "2"
	// 		s.Emit("reply", "You Are connected to chat room "+roomId)
	// 		fmt.Println("A user connected:", s.ID(), "to Room", roomId)
	// 	} else {
	// 		s.Emit("reply", "Welcone to my chat room")
	// 	}
	// 	// fmt.Println("User joined room 2:", s.ID())

	// 	return nil
	// })

	// // Handle a "msg" event and broadcast a reply to the sender
	// server.OnEvent("/chat", "set_user", func(s socketio.Conn, name string) {
	// 	user := User{
	// 		name: name,
	// 	}
	// 	// Associate the user struct with the connection ID
	// 	users.Store(s.ID(), user)
	// 	fmt.Printf("User registered: %+v with Socket ID %s\n", user, s.ID())
	// 	s.Emit("set_user_success", "User registered successfully!")
	// })

	// server.OnEvent("/chat", "msg", func(s socketio.Conn, roomId string, msg string) {
	// 	// Print the received message
	// 	fmt.Println("RoomId is:", roomId, "Received message:", msg)
	// 	user, ok := users.Load(s.ID())
	// 	var new_message = ""
	// 	if ok {
	// 		new_message = user.(User).name + " :- " + msg
	// 	} else {
	// 		new_message = "Anonymous User :- " + msg
	// 	}

	// 	// Broadcast to all clients in the room with ID "2"
	// 	server.BroadcastToRoom("/chat", roomId, "reply", new_message)
	// })

	// // server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) {
	// // 	// s.SetContext(msg)
	// // 	// s.Join("chat")
	// // 	fmt.Println("msg:", msg)
	// // 	s.Emit("reply", "have "+msg)
	// // 	// return "recv " + msg
	// // })

	// server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
	// 	fmt.Println("notice:", msg)
	// 	s.Emit("reply", "have "+msg)
	// })

	// server.OnEvent("/", "bye", func(s socketio.Conn) string {
	// 	last := s.Context().(string)
	// 	s.Emit("bye", last)
	// 	s.Close()
	// 	return last
	// })

	// go server.Serve()
	// defer server.Close()

	// http.Handle("/socket.io/", server)
	// http.Handle("/", http.FileServer(http.Dir("./static")))
	// log.Println("Serving at localhost:5001...")
	// log.Fatal(http.ListenAndServe(":5001", nil))

}
