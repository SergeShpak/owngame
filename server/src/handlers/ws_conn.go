package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SergeyShpak/owngame/server/src/model"
	"github.com/SergeyShpak/owngame/server/src/types"
	"github.com/SergeyShpak/owngame/server/src/ws"
)

func RoomJoin(dl *model.DataLayer) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.RoomJoinRequest
		c.BindJSON(&req)
		if err := dl.Rooms.CheckPassword(req.RoomName, req.Password); err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		role, err := dl.Rooms.JoinRoom(req.RoomName, req.Login)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		client, err := ws.UpgradeConnection(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err := client.WriteMsg(fmt.Sprintf("Hello from ws! Your role is: %v", role)); err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err := client.Close(); err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}
