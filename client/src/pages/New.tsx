import * as React from "react";
import localStorage from "../lib/localStorage";

export const New: React.StatelessComponent<{}> = props => {
  let token = localStorage.get("ws-token");
  if (token == null) {
    return <h1>No token found</h1>;
  }
  return <h1>New room! {token}</h1>;
};
