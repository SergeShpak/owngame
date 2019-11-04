export interface CreateRoomRequest {
  roomName: string;
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

enum ParticipantRole {
  Host = 1,
  Participant
}

export interface Participant {
  login: string;
  role: ParticipantRole;
}

export const WS_MSG_TYPE_PARTICIPANTS = "participants";

export interface WSParticipants {
  participants: Participant[];
}
