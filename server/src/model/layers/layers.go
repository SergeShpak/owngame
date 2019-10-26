package layers

import "github.com/SergeyShpak/owngame/server/src/types"

type RoomsDataLayer interface {
	CreateRoom(r *types.RoomCreateRequest) error
}
