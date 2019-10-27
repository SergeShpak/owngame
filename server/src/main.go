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
	api := r.Group("/api/v1")

	api.POST("/room", handlers.RoomCreate(dl))
	api.PUT("/room/:roomName", handlers.RoomJoin(dl))
	r.Run(":8080")
}
