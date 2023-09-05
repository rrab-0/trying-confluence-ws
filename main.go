package main

import (
	"gin-ws-kafka/myws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	myws.DoWriter(myws.WS, "hello, websocket!")

	c.JSON(http.StatusOK, gin.H{
		"message": "Service is up and running.",
	})
}

func main() {
	app := gin.Default()

	app.GET("/ws", myws.Listener)
	app.GET("/", Homepage)

	app.Run("localhost:8080")
}
