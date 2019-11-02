package layers

import (
	"github.com/SergeyShpak/owngame/server/src/types"
	"github.com/SergeyShpak/owngame/server/src/ws"
)

type RoomsDataLayer interface {
	CreateRoom(r *types.RoomCreateRequest, hostToken string) error
	CheckPassword(roomName string, password string) error
	JoinRoom(roomName string, login string) (types.PlayerRole, error)
	GetParticipants(roomName string) ([]string, error)
}

type WebsocketConnectionLayer interface {
	PrepareConnection(token string, connMeta *types.ConnectionMeta) error
	GetConnectionMeta(token string) (*types.ConnectionMeta, error)
	EstablishConnection(c *ws.Client, connMeta *types.ConnectionMeta) error
	GetConnection(connMeta *types.ConnectionMeta) (*ws.Client, error)
}
