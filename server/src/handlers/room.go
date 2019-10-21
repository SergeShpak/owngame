package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RoomCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Login    string `json:"login" binding:"required"`
}

func RoomCreate(c *gin.Context) {
	var req RoomCreateRequest
	c.BindJSON(&req)
	fmt.Println("From create", req)
}
