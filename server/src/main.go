package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/SergeyShpak/owngame/server/src/handlers"
	"github.com/SergeyShpak/owngame/server/src/model"
)

func main() {
	dl, err := model.NewDataLayer()
	if err != nil {
		panic("failed to initialize the data layer")
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/room", handlers.RoomCreate(dl))
	r.Run(":8080")
}
