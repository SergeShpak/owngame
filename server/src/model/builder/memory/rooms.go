package memory

import (
	"fmt"

	"github.com/SergeyShpak/owngame/server/src/model/layers"
	"github.com/SergeyShpak/owngame/server/src/types"
)

type memoryRoomLayer struct {
	rooms       *rooms
	roomPlayers *roomPlayers
}

func NewMemoryRoomLayer() (layers.RoomsDataLayer, error) {
	m := &memoryRoomLayer{
		rooms:       newRooms(),
		roomPlayers: newRoomPlayers(),
	}
	return m, nil
}

func (m *memoryRoomLayer) CreateRoom(r *types.RoomCreateRequest, roomToken string) error {
	rMeta := &roomMeta{
		Name:            r.RoomName,
		Password:        r.Password,
		MaxPlayersCount: 3,
	}
	if ok := m.rooms.PutRoomMeta(rMeta); !ok {
		return fmt.Errorf("room %s already exists", r.RoomName)
	}
	return nil
}

func (m *memoryRoomLayer) CheckPassword(roomName string, password string) error {
	meta, ok := m.rooms.GetRoomMeta(roomName)
	if !ok {
		return fmt.Errorf("rooms %s not found", roomName)
	}
	if meta.Password != password {
		return fmt.Errorf("password validation for room %s failed, expected: %s, actual: %s", roomName, meta.Password, password)
	}
	return nil
}

func (m *memoryRoomLayer) JoinRoom(roomName string, login string) (types.PlayerRole, error) {
	meta, ok := m.rooms.GetRoomMeta(roomName)
	if !ok {
		return types.PlayerRoleParticipant, fmt.Errorf("room %s not found", roomName)
	}
	if ok := m.roomPlayers.PutHost(roomName, login); ok {
		return types.PlayerRoleHost, nil
	}
	if ok := m.roomPlayers.AddParticipant(meta, login); ok {
		return types.PlayerRoleParticipant, nil
	}
	return types.PlayerRoleParticipant, fmt.Errorf("failed to join the room %s", roomName)
}

func (m *memoryRoomLayer) GetParticipants(roomName string) ([]string, error) {
	p, ok := m.roomPlayers.GetPlayers(roomName)
	if !ok {
		return nil, fmt.Errorf("room %s not found", roomName)
	}
	playersCount := len(p.Players) + 1
	ps := make([]string, playersCount)
	for i, player := range p.Players {
		ps[i] = player
	}
	ps[len(p.Players)] = p.Host
	return ps, nil
}

type rooms keyValStore

func newRooms() *rooms {
	r := (rooms)(*newKeyValStore())
	return &r
}

type roomMeta struct {
	Name            string
	Password        string
	MaxPlayersCount int
}

func (r *rooms) PutRoomMeta(meta *roomMeta) bool {
	kvs := (*keyValStore)(r)
	roomName := meta.Name
	return kvs.Put(roomName, meta)
}

func (r *rooms) GetRoomMeta(roomName string) (*roomMeta, bool) {
	kvs := (*keyValStore)(r)
	roomMetaIface, ok := kvs.Get(roomName)
	if !ok {
		return nil, false
	}
	roomMeta := (roomMetaIface).(*roomMeta)
	return roomMeta, true
}

type roomPlayers keyValStore

func newRoomPlayers() *roomPlayers {
	rp := (roomPlayers)(*newKeyValStore())
	return &rp
}

type players struct {
	Host      string
	Players   []string
	Observers []string
}

func newPlayers() *players {
	p := &players{
		Players:   make([]string, 0, 3),
		Observers: make([]string, 0),
	}
	return p
}

func (rp *roomPlayers) PutHost(roomName string, login string) bool {
	kvs := (*keyValStore)(rp)
	ok := kvs.Alter(roomName, func(playersIface interface{}, exist bool) (interface{}, bool) {
		if !exist {
			p := newPlayers()
			p.Host = login
			return p, true
		}
		p := playersIface.(*players)
		if len(p.Host) != 0 {
			return nil, false
		}
		p.Host = login
		return p, true
	})
	return ok
}

func (rp *roomPlayers) AddParticipant(meta *roomMeta, login string) bool {
	kvs := (*keyValStore)(rp)
	ok := kvs.Alter(meta.Name, func(playersIface interface{}, exist bool) (interface{}, bool) {
		if !exist {
			return nil, false
		}
		p := playersIface.(*players)
		if meta.MaxPlayersCount == len(p.Players) {
			return nil, false
		}
		p.Players = append(p.Players, login)
		return p, true
	})
	return ok
}

func (rp *roomPlayers) GetPlayers(roomName string) (*players, bool) {
	kvs := (*keyValStore)(rp)
	pIface, ok := kvs.Get(roomName)
	if !ok {
		return nil, false
	}
	players := pIface.(*players)
	return players, true
}

type roomAdmin keyValStore

func newRoomAdmin() *roomAdmin {
	ra := (roomAdmin)(*newKeyValStore())
	return &ra
}

func (ra *roomAdmin) PutAdmin(roomName string, adminToken string) bool {
	kvs := (*keyValStore)(ra)
	ok := kvs.Alter(roomName, func(tokensIface interface{}, exist bool) (interface{}, bool) {
		if !exist {
			t := make([]string, 1)
			t[0] = adminToken
			return t, true
		}
		t := tokensIface.([]string)
		for _, at := range t {
			if adminToken == at {
				return t, true
			}
		}
		t = append(t, adminToken)
		return t, true
	})
	return ok
}

func (ra *roomAdmin) DeleteAdmin(roomName string, adminToken string) bool {
	t, ok := ra.GetAdmins(roomName)
	if !ok {
		return false
	}
	for i, at := range t {
		if adminToken == at {
			t = append(t[:i], t[i+1:]...)
		}
	}
	return true
}

func (ra *roomAdmin) GetAdmins(roomName string) ([]string, bool) {
	kvs := (*keyValStore)(ra)
	tokensIface, ok := kvs.Get(roomName)
	if !ok {
		return nil, false
	}
	t := tokensIface.([]string)
	return t, true
}
