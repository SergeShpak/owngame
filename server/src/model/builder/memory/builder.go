package memory

import (
	"github.com/SergeyShpak/owngame/server/src/model/builder/ibuilder"
	"github.com/SergeyShpak/owngame/server/src/model/layers"
)

func NewMemoryDataLayerBuilder() (ibuilder.DataLayerBuilder, error) {
	m := &memoryDataLayerBuilder{}
	return m, nil
}

type memoryDataLayerBuilder struct{}

func (m *memoryDataLayerBuilder) BuildRoomLayer() (layers.RoomsDataLayer, error) {
	l, err := NewMemoryRoomLayer()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (m *memoryDataLayerBuilder) BuildWebsocketConnectionLayer() (layers.WebsocketConnectionLayer, error) {
	ws, err := NewMemoryWebsocketConnectionLayer()
	if err != nil {
		return nil, err
	}
	return ws, nil
}
