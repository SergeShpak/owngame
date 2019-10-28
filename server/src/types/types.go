package types

type RoomCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Login    string `json:"login" binding:"required"`
}

type RoomCreateResponse struct {
	Token string `json:"token" binding:"required"`
}

type RoomJoinRequest struct {
	Login    string `json:"login" binding:"required"`
	RoomName string `json:"roomName" binding: "required"`
	Password string `json:"password" binding: "required"`
}

type RoomJoinResponse struct {
	Token string `json:"token"`
}

type PlayerRole int

const (
	PlayerRoleHost PlayerRole = iota + 1
	PlayerRoleParticipant
)
