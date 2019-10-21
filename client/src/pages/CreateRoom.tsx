import * as React from "react";

import { CreateRoomForm } from "../components/CreateRoomForm";

export const CreateRoom: React.StatelessComponent<{}> = props => {
  return (
    <div className="col-md-6 offset-md-3">
      <CreateRoomForm />
    </div>
  );
};
