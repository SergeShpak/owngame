package memory

import (
	"github.com/SergeyShpak/owngame/server/src/model/layers"
	"github.com/SergeyShpak/owngame/server/src/types"
)

func NewMemoryRoomLayer() (layers.RoomsDataLayer, error) {
	m := &memoryRoomLayer{}
	return m, nil
}

type memoryRoomLayer struct{}

func (m *memoryRoomLayer) CreateRoom(r *types.RoomCreateRequest) error {
	return nil
}
