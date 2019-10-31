import * as React from "react";

import * as Paths from "../paths";
import { CreateRoomForm } from "../components/CreateRoomForm";
import { Redirect } from "react-router-dom";

interface Props {
  history?: any;
}

export const CreateRoom: React.StatelessComponent<Props> = props => {
  const [create, setCreate] = React.useState(false);
  const [token, setToken] = React.useState("");
  const onCreate = () => {
    setCreate(true);
  };
  if (create) {
    return (
      <Redirect
        push
        to={{
          pathname: Paths.ROOM,
          state: {
            token
          }
        }}
      />
    );
  }
  return (
    <div>
      <div className="col-md-6 offset-md-3">
        <CreateRoomForm onCreate={onCreate} setToken={setToken} />
      </div>
    </div>
  );
};
