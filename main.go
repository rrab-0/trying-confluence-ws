package main

import (
	"gin-ws-kafka/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	websocket.DoWriter("hello, websocket!")

	c.JSON(http.StatusOK, gin.H{
		"message": "Service is up and running.",
	})
}

type MsgFromKafka struct {
	Message string `form:"msg"`
}

func main() {
	app := gin.Default()

	app.GET("/ws", websocket.Listener)
	app.GET("/", Homepage)

	app.Run("localhost:8080")
}
