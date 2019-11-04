package handlers

import (
	"github.com/SergeyShpak/owngame/server/src/model"
	"github.com/SergeyShpak/owngame/server/src/types"
)

func prepareConnection(dl *model.DataLayer, token string, roomName string, login string) error {
	connMeta := &types.ConnectionMeta{
		RoomName: roomName,
		Login:    login,
	}
	if err := dl.WebsocketConnection.PrepareConnection(token, connMeta); err != nil {
		return err
	}
	return nil
}
