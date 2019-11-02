package memory

import (
	"fmt"

	"github.com/SergeyShpak/owngame/server/src/model/layers"
	"github.com/SergeyShpak/owngame/server/src/types"
	"github.com/SergeyShpak/owngame/server/src/ws"
)

type memoryWebsocketConnectionLayer struct {
	tokenStore *wsTokenStore
	connStore  *wsConnectionStore
}

func NewMemoryWebsocketConnectionLayer() (layers.WebsocketConnectionLayer, error) {
	ws := &memoryWebsocketConnectionLayer{
		tokenStore: newWSTokenStore(),
		connStore:  newWSConnectionStore(),
	}
	return ws, nil
}

func (ws *memoryWebsocketConnectionLayer) PrepareConnection(token string, connMeta *types.ConnectionMeta) error {
	ws.tokenStore.PutConnectionToken(token, connMeta)
	return nil
}

func (ws *memoryWebsocketConnectionLayer) GetConnectionMeta(token string) (*types.ConnectionMeta, error) {
	connMeta, err := ws.tokenStore.GetRoomInfo(token)
	if err != nil {
		return nil, err
	}
	return connMeta, nil
}

func (ws *memoryWebsocketConnectionLayer) EstablishConnection(c *ws.Client, connMeta *types.ConnectionMeta) error {
	if err := ws.connStore.PutConnection(c, connMeta); err != nil {
		return err
	}
	return nil
}

func (ws *memoryWebsocketConnectionLayer) GetConnection(connMeta *types.ConnectionMeta) (*ws.Client, error) {
	c, err := ws.connStore.GetConnection(connMeta)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type wsTokenStore keyValStore

func newWSTokenStore() *wsTokenStore {
	s := (wsTokenStore)(*newKeyValStore())
	return &s
}

func (ws *wsTokenStore) PutConnectionToken(t string, connMeta *types.ConnectionMeta) {
	kvs := (*keyValStore)(ws)
	kvs.Set(t, connMeta)
}

func (ws *wsTokenStore) GetRoomInfo(t string) (*types.ConnectionMeta, error) {
	kvs := (*keyValStore)(ws)
	roomInfoIface, ok := kvs.Get(t)
	if !ok {
		return nil, fmt.Errorf("token %s was not found", t)
	}
	i := roomInfoIface.(*types.ConnectionMeta)
	return i, nil
}

type wsConnectionStore keyValStore

func newWSConnectionStore() *wsConnectionStore {
	s := (wsConnectionStore)(*newKeyValStore())
	return &s
}

func (s *wsConnectionStore) PutConnection(c *ws.Client, connMeta *types.ConnectionMeta) error {
	kvs := (*keyValStore)(s)
	key, err := s.getKeyFromConnMeta(connMeta)
	if err != nil {
		return err
	}
	kvs.Set(key, c)
	return nil
}

func (s *wsConnectionStore) GetConnection(connMeta *types.ConnectionMeta) (*ws.Client, error) {
	kvs := (*keyValStore)(s)
	key, err := s.getKeyFromConnMeta(connMeta)
	if err != nil {
		return nil, err
	}
	cIface, ok := kvs.Get(key)
	if !ok {
		return nil, fmt.Errorf("connection %s not found", key)
	}
	c := cIface.(*ws.Client)
	return c, nil
}

func (s *wsConnectionStore) getKeyFromConnMeta(connMeta *types.ConnectionMeta) (string, error) {
	if connMeta == nil {
		return "", fmt.Errorf("connection meta is nil")
	}
	if len(connMeta.RoomName) == 0 || len(connMeta.Login) == 0 {
		return "", fmt.Errorf("connection meta is invalid")
	}
	key := fmt.Sprintf("%s/%s", connMeta.RoomName, connMeta.Login)
	return key, nil
}
