package myws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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

	go reader(WS)
}

func reader(conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func writer(conn *websocket.Conn, message string) error {
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		log.Println("Error sending message to websocket:", err)
		return err
	}

	return nil
}

func DoWriter(conn *websocket.Conn, message string) {
	go writer(conn, message)
}
