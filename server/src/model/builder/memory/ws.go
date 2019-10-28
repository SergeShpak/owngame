package memory

import (
	"fmt"

	"github.com/SergeyShpak/owngame/server/src/model/layers"
)

type memoryWebsocketConnectionLayer struct {
	wsConnectionStore *wsConnectionStore
}

func NewMemoryWebsocketConnectionLayer() (layers.WebsocketConnectionLayer, error) {
	ws := &memoryWebsocketConnectionLayer{
		wsConnectionStore: newWSConnectionStore(),
	}
	return ws, nil
}

func (ws *memoryWebsocketConnectionLayer) PrepareConnection(token string, roomName string, login string) error {
	ws.wsConnectionStore.PutConnectionToken(token, &roomInfo{
		RoomName: roomName,
		Login:    login,
	})
	return nil
}

func (ws *memoryWebsocketConnectionLayer) EstablishConnection(token string) (string, string, error) {
	info, err := ws.wsConnectionStore.RemoveConnectionToken(token)
	if err != nil {
		return "", "", err
	}
	return info.RoomName, info.Login, nil
}

type wsConnectionStore keyValStore

func newWSConnectionStore() *wsConnectionStore {
	ws := (wsConnectionStore)(*newKeyValStore())
	return &ws
}

type roomInfo struct {
	RoomName string
	Login    string
}

func (ws *wsConnectionStore) PutConnectionToken(t string, i *roomInfo) {
	kvs := (*keyValStore)(ws)
	kvs.Set(t, i)
}

func (ws *wsConnectionStore) RemoveConnectionToken(t string) (*roomInfo, error) {
	kvs := (*keyValStore)(ws)
	roomInfoIface, ok := kvs.Pop(t)
	if !ok {
		return nil, fmt.Errorf("token %s was not found", t)
	}
	i := roomInfoIface.(*roomInfo)
	return i, nil
}
