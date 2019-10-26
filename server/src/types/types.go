package types

type RoomCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Login    string `json:"login" binding:"required"`
}

type RoomCreateResponse struct {
	Token string `json:"token" binding:"required"`
}
