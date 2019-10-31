import * as React from "react";
import { WS } from "../lib/ws";
import { createHook } from "async_hooks";

interface RoomProps {
  token: string;
}

export const Room: React.StatelessComponent<RoomProps> = props => {
  const [connStatus, setConnStatus] = React.useState("loading");

  if (props.token == null || props.token.length === 0) {
    return <h1>No token found</h1>;
  }
  if (connStatus == "loading") {
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
      return <h1>New room! {props.token}</h1>;
    case "failed":
      return <h1>Connection failed</h1>;
    case "loading":
      return <h1>Loading...</h1>;
    default:
      throw `state ${connStatus} is unknown`;
  }
};

function onWSMessage(data: MessageEvent): void {
  console.log(data);
}
