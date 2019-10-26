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
		if err := model.Rooms.CreateRoom(&req); err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("an error occurred"))
			return
		}
	}
}
