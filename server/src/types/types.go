package types

type RoomCreateRequest struct {
	RoomName string `json:"roomName" binding:"required"`
	Password string `json:"password" binding:"required"`
	Login    string `json:"login" binding:"required"`
}

type RoomCreateResponse struct {
	Token string `json:"token" binding:"required"`
}

type RoomJoinRequest RoomCreateRequest

type RoomJoinResponse struct {
	Token string `json:"token"`
}

type PlayerRole int

const (
	PlayerRoleHost PlayerRole = iota + 1
	PlayerRoleParticipant
)

type ConnectionMeta struct {
	RoomName string
	Login    string
}

type Participant struct {
	Login string     `json:"login"`
	Role  PlayerRole `json:"role"`
}

type WSMessageType string

const (
	WSMessageTypeParticipants = "participants"
)

type WSMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type WSMsgParticipant struct {
	Participants []Participant `json:"participants"`
}
