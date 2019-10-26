package memory

import (
	"fmt"
	"sync"

	"github.com/SergeyShpak/owngame/server/src/model/layers"
	"github.com/SergeyShpak/owngame/server/src/types"
)

func NewMemoryRoomLayer() (layers.RoomsDataLayer, error) {
	m := &memoryRoomLayer{
		rooms: newRooms(),
	}
	return m, nil
}

type memoryRoomLayer struct {
	rooms   *rooms
	players *players
}

func (m *memoryRoomLayer) CreateRoom(r *types.RoomCreateRequest, roomToken string) error {
	rMeta := &roomMeta{
		Name:            r.Name,
		Password:        r.Password,
		MaxPlayersCount: 3,
		Token:           roomToken,
	}
	if ok := m.rooms.Set(rMeta); !ok {
		return fmt.Errorf("room %s already exists", r.Name)
	}
	return nil
}

type rooms struct {
	mux  sync.RWMutex
	vals map[string]*roomMeta
}

type roomMeta struct {
	Token           string
	Name            string
	Password        string
	MaxPlayersCount int
}

func newRooms() *rooms {
	r := &rooms{
		mux:  sync.RWMutex{},
		vals: make(map[string]*roomMeta),
	}
	return r
}

func (r *rooms) GetMeta(roomName string) (*roomMeta, bool) {
	r.mux.RLock()
	meta, ok := r.vals[roomName]
	r.mux.RUnlock()
	return meta, ok
}

func (r *rooms) Set(meta *roomMeta) bool {
	if meta == nil || len(meta.Name) == 0 {
		return false
	}
	r.mux.Lock()
	if _, ok := r.vals[meta.Name]; ok {
		r.mux.Unlock()
		return false
	}
	r.vals[meta.Name] = meta
	r.mux.Unlock()
	return true
}

type players struct {
	mux  sync.RWMutex
	vals map[string]*roomPlayers
}

func newPlayers() *players {
	p := &players{
		mux:  sync.RWMutex{},
		vals: make(map[string]*roomPlayers),
	}
	return p
}

type roomPlayers struct {
	Host    string
	Players []string
}

func (p *players) AddHost(meta *roomMeta, hostNickname string) bool {
	p.mux.Lock()
	if _, ok := p.vals[meta.Name]; ok {
		p.mux.Unlock()
		return false
	}
	p.vals[meta.Name] = &roomPlayers{
		Host:    hostNickname,
		Players: make([]string, 0, meta.MaxPlayersCount),
	}
	p.mux.Unlock()
	return true
}
