export interface CreateRoomRequest {
  name: string;
  password: string;
  login: string;
}

export interface CreateRoomResponse {
  token: string;
}

export interface RoomJoinRequest {
  login: string;
  roomName: string;
  password: string;
}

export interface RoomJoinResponse {
  token: string;
}
