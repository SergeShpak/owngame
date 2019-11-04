import * as React from "react";
import { WS } from "../lib/ws";
import {
  WebsocketMessage,
  WS_MSG_TYPE_PARTICIPANTS,
  WSParticipants,
  Participant
} from "../lib/types";
import { Room as RoomComponent } from "../components/Room";

interface RoomProps {
  token: string;
}

export const Room: React.StatelessComponent<RoomProps> = props => {
  const [connStatus, setConnStatus] = React.useState("loading");
  const [participants, setParticipants] = React.useState<Participant[]>([]);

  const onWSMessage = (data: MessageEvent) => {
    console.log("Received msg: ", data.data);
    const msg: WebsocketMessage = JSON.parse(data.data);
    switch (msg.type) {
      case WS_MSG_TYPE_PARTICIPANTS:
        const participants: WSParticipants = JSON.parse(atob(msg.message));
        updateParticipants(participants);
        break;
      default:
        throw new Error(`ws message type ${msg.type} is unknown`);
    }
  };

  const updateParticipants = (msg: WSParticipants) => {
    setParticipants(msg.participants);
  };

  if (props.token == null || props.token.length === 0) {
    return <h1>No token found</h1>;
  }
  if (connStatus === "loading") {
    let conn = new WS();
    conn
      .open("/api/v1/room/ws", onWSMessage, [
        { key: "token", value: props.token }
      ])
      .then(() => {
        setConnStatus("succeeded");
      })
      .catch(e => {
        console.log(e);
        setConnStatus("failed");
      });
  }
  switch (connStatus) {
    case "succeeded":
      return <RoomComponent participants={participants} />;
    case "failed":
      return <h1>Connection failed</h1>;
    case "loading":
      return <h1>Loading...</h1>;
    default:
      throw new Error(`state ${connStatus} is unknown`);
  }
};
