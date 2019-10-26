package builder

import (
	"github.com/SergeyShpak/owngame/server/src/model/builder/ibuilder"
	"github.com/SergeyShpak/owngame/server/src/model/builder/memory"
)

func NewDataLayerBuilder() (ibuilder.DataLayerBuilder, error) {
	b, err := memory.NewMemoryDataLayerBuilder()
	if err != nil {
		return nil, err
	}
	return b, nil
}
