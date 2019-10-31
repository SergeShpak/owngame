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

export interface WebsocketMessage {
  type: string;
  message: string;
}

export const WS_MSG_TYPE_PARTICIPANTS = "participants";

export interface WSParticipants {
  logins: string[];
}
