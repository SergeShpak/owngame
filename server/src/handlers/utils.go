package handlers

import (
	"encoding/base64"
	"fmt"

	"github.com/SergeyShpak/owngame/server/src/utils"
)

func generateRoomToken(roomName string, login string) (string, error) {
	nonce, err := utils.GenerateToken(16)
	if err != nil {
		return "", err
	}
	roomName64 := base64.RawURLEncoding.EncodeToString([]byte(roomName))
	login64 := base64.RawURLEncoding.EncodeToString([]byte(login))
	token := fmt.Sprintf("%s.%s.%s", nonce, roomName64, login64)
	return token, nil
}
