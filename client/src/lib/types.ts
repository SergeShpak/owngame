export interface CreateRoomRequest {
  name: string;
  password: string;
  login: string;
}

export interface CreateRoomResponse {
  token: string;
}
