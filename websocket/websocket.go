package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// var (
// 	upgrader = websocket.Upgrader{
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true
// 		},
// 	}
// 	clients   = make(map[*websocket.Conn]bool)
// 	clientsMu sync.Mutex
// )

// func Listener(c *gin.Context) {
// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Println("Error upgrading connection:", err)
// 		return
// 	}
// 	defer func() {
// 		conn.Close()

// 		clientsMu.Lock()
// 		delete(clients, conn)
// 		clientsMu.Unlock()
// 	}()

// 	clientsMu.Lock()
// 	clients[conn] = true
// 	clientsMu.Unlock()

// 	for {
// 		_, _, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("Error reading message:", err)
// 			break
// 		}
// 	}
// }

// func BroadcastMessage(message string) {
// 	clientsMu.Lock()
// 	defer clientsMu.Unlock()

// 	for conn := range clients {
// 		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
// 		if err != nil {
// 			log.Println("Error sending message to client:", err)
// 		}
// 	}
// }

// func SendMessage(message string) {
// 	go BroadcastMessage(message)
// }

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var WS *websocket.Conn

func Listener(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	WS = ws
	if err != nil {
		log.Println(err)
	}

	reader(WS)
}

func reader(conn *websocket.Conn) {
	for {
		// messageType, p, err := conn.ReadMessage()
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// log.Println(string(p))

		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}

func writer(message string) error {
	if err := WS.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		log.Println("Error sending message to websocket:", err)
		return err
	}

	log.Println("👍", WS, "👍")
	return nil
}

func DoWriter(message string) {
	go writer(message)
}
