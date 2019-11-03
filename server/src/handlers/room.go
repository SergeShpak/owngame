package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SergeyShpak/owngame/server/src/model"
	"github.com/SergeyShpak/owngame/server/src/types"
)

func RoomCreate(model *model.DataLayer) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.RoomCreateRequest
		c.BindJSON(&req)
		roomToken, err := generateRoomToken(req.RoomName, req.Login)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("an error occurred"))
		}
		if err := model.Rooms.CreateRoom(&req, roomToken); err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("an error occurred"))
			return
		}
		resp := types.RoomCreateResponse{
			Token: roomToken,
		}
		c.JSON(http.StatusCreated, resp)
	}
}
