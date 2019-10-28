package layers

import "github.com/SergeyShpak/owngame/server/src/types"

type RoomsDataLayer interface {
	CreateRoom(r *types.RoomCreateRequest, hostToken string) error
	CheckPassword(roomName string, password string) error
	JoinRoom(roomName string, login string) (types.PlayerRole, error)
}

type WebsocketConnectionLayer interface {
	PrepareConnection(token string, roomName string, login string) error
	EstablishConnection(token string) (string, string, error)
}
