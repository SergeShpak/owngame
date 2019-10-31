package handlers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SergeyShpak/owngame/server/src/model"
	"github.com/SergeyShpak/owngame/server/src/types"
	"github.com/SergeyShpak/owngame/server/src/utils"
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
		_, err := dl.Rooms.JoinRoom(req.RoomName, req.Login)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		token, err := generateWebsocketToken(req.RoomName, req.Login)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err := dl.WebsocketConnection.PrepareConnection(token, req.RoomName, req.Login); err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		resp := types.RoomJoinResponse{
			Token: token,
		}
		c.JSON(http.StatusOK, resp)
	}
}

func RoomCreateWSConn(dl *model.DataLayer) func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		_, login, err := dl.WebsocketConnection.EstablishConnection(token)
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
		wsMsg, err := ws.NewMsgParticipants([]string{login})
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		client.WriteMsg(wsMsg)
	}
}

func generateWebsocketToken(roomName string, login string) (string, error) {
	nonce, err := utils.GenerateToken(16)
	if err != nil {
		return "", err
	}
	roomName64 := base64.RawURLEncoding.EncodeToString([]byte(roomName))
	login64 := base64.RawURLEncoding.EncodeToString([]byte(login))
	token := fmt.Sprintf("%s.%s.%s", nonce, roomName64, login64)
	return token, nil
}
