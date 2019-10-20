package main

import (
	"github.com/gin-gonic/gin"

	"github.com/SergeyShpak/owngame/server/src/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/room", handlers.RoomCreate)
	r.Run(":8080")
}
