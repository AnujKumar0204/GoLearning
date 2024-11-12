package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	socketio "github.com/googollee/go-socket.io"
)

func main() {

	// db, err := database.ConnectDB()

	// if err != nil {
	// 	fmt.Println("Failed to connect to database: %v", err)
	// 	return
	// }

	// app := fiber.New()

	// router.RouteHandler(app, db)

	// app.Listen(":8080")

	// app := fiber.New()

	// app.Use("/ws", func(c *fiber.Ctx) error {
	// 	// IsWebSocketUpgrade returns true if the client
	// 	// requested upgrade to the WebSocket protocol.
	// 	if websocket.IsWebSocketUpgrade(c) {
	// 		c.Locals("allowed", true)
	// 		return c.Next()
	// 	}
	// 	return fiber.ErrUpgradeRequired
	// })

	// app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
	// 	// c.Locals is added to the *websocket.Conn
	// 	log.Println(c.Locals("allowed"))  // true
	// 	log.Println(c.Params("id"))       // 123
	// 	log.Println(c.Query("v"))         // 1.0
	// 	log.Println(c.Cookies("session")) // ""

	// 	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	// 	var (
	// 		mt  int
	// 		msg []byte
	// 		err error
	// 	)
	// 	for {
	// 		if mt, msg, err = c.ReadMessage(); err != nil {
	// 			log.Println("read:", mt, err)
	// 			break
	// 		}
	// 		log.Printf("recv:", mt)
	// 		log.Printf("recv: %s", msg)

	// 		if err = c.WriteMessage(mt, msg); err != nil {
	// 			log.Println("write:", err)
	// 			break
	// 		}
	// 	}

	// }))

	// log.Fatal(app.Listen(":3000"))

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	// Handle a "msg" event and broadcast a reply to the sender
	// server.OnEvent("/chat", "joinRoom", func(s socketio.Conn, roomId string) {

	// 	// The client sent a roomId, we can join the user to that room
	// 	s.Join(roomId) // Join the requested room
	// 	s.Emit("reply", "have have connected to room "+roomId)
	// 	fmt.Println("User wants to join room:", roomId)
	// })

	// Handle a "msg" event and broadcast a reply to the sender
	server.OnEvent("/chat", "msg", func(s socketio.Conn, roomId string, msg string) {
		// Print the received message
		fmt.Println("RoomId is:", roomId, "Received message:", msg)
		// fmt.Println("RoomId is:", roomId)
		// fmt.Println("Received message:", s.Rooms())

		// Emit a "reply" event to the sender
		// s.Emit("reply", "have "+msg)

		// // You can also broadcast to all users in room 2
		// s.Emit("reply", "Message received: "+msg)

		// Broadcast to all clients in the room with ID "2"
		server.BroadcastToRoom("/chat", roomId, "reply", msg)
	})

	// Handle user connections
	server.OnConnect("/chat", func(s socketio.Conn) error {
		// fmt.Println("A user connected:", s.ID())

		queryString := s.URL().RawQuery

		// Parse the query string
		values, err := url.ParseQuery(queryString)
		if err != nil {
			log.Fatal(err)
		}

		// Get the 'roomId' value
		roomId := values.Get("roomId")
		// fmt.Println("RoomId is:", roomId)

		// Automatically add the user to room 2 upon connection
		if roomId != "" {
			s.Join(roomId) // The user is automatically added to room "2"
			s.Emit("reply", "You Are connected to chat room"+roomId)
			fmt.Println("A user connected:", s.ID(), "to Room", roomId)
		} else {
			s.Emit("reply", "Welcone to my chat room")
		}
		// fmt.Println("User joined room 2:", s.ID())

		return nil
	})

	// server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) {
	// 	// s.SetContext(msg)
	// 	// s.Join("chat")
	// 	fmt.Println("msg:", msg)
	// 	s.Emit("reply", "have "+msg)
	// 	// return "recv " + msg
	// })

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		// server.Remove(s.ID())
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Add the Remove session id. Fixed the connection & mem leak
		// server.Remove(s.ID())
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Serving at localhost:5001...")
	log.Fatal(http.ListenAndServe(":5001", nil))

}
