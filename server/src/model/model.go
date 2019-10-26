package model

import (
	"github.com/SergeyShpak/owngame/server/src/model/builder"
	"github.com/SergeyShpak/owngame/server/src/model/layers"
)

type DataLayer struct {
	Rooms layers.RoomsDataLayer
}

func NewDataLayer() (*DataLayer, error) {
	b, err := builder.NewDataLayerBuilder()
	if err != nil {
		return nil, err
	}
	dl := &DataLayer{}
	dl.Rooms, err = b.BuildRoomLayer()
	if err != nil {
		return nil, err
	}
	return dl, nil
}
