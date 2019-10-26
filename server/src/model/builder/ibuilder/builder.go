package ibuilder

import (
	"github.com/SergeyShpak/owngame/server/src/model/layers"
)

type DataLayerBuilder interface {
	BuildRoomLayer() (layers.RoomsDataLayer, error)
}
