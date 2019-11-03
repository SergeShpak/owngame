package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SergeyShpak/owngame/server/src/model"
	"github.com/SergeyShpak/owngame/server/src/types"
	"github.com/SergeyShpak/owngame/server/src/ws"
)

type WsConn struct {
	dl *model.DataLayer
}

func NewWsConn(dl *model.DataLayer) *WsConn {
	wsConn := &WsConn{
		dl: dl,
	}
	return wsConn
}

func (conn *WsConn) RoomJoin() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.RoomJoinRequest
		c.BindJSON(&req)
		if err := conn.dl.Rooms.CheckPassword(req.RoomName, req.Password); err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		_, err := conn.dl.Rooms.JoinRoom(req.RoomName, req.Login)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		token, err := generateRoomToken(req.RoomName, req.Login)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		connMeta := &types.ConnectionMeta{
			RoomName: req.RoomName,
			Login:    req.Login,
		}
		if err := conn.dl.WebsocketConnection.PrepareConnection(token, connMeta); err != nil {
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

func (conn *WsConn) RoomCreateWSConn() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		connMeta, err := conn.dl.WebsocketConnection.GetConnectionMeta(token)
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
		conn.dl.WebsocketConnection.EstablishConnection(client, connMeta)
		if err := conn.BroadcastParticipantList(connMeta.RoomName); err != nil {
			log.Printf("[ERROR]: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}

func (conn *WsConn) BroadcastParticipantList(roomName string) error {
	participants, err := conn.dl.Rooms.GetParticipants(roomName)
	if err != nil {
		return err
	}
	wsMsg, err := ws.NewMsgParticipants(participants)
	if err != nil {
		return err
	}
	for _, p := range participants {
		connMeta := &types.ConnectionMeta{
			RoomName: roomName,
			Login:    p,
		}
		c, err := conn.dl.WebsocketConnection.GetConnection(connMeta)
		if err != nil {
			// TODO(SSH): add logging
		}
		c.WriteMsg(wsMsg)
	}
	return nil
}
