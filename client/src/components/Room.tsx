import * as React from "react";
import { Participant } from "../lib/types";

interface RoomProps {
  participants: Participant[];
}

export const Room: React.StatelessComponent<RoomProps> = props => {
  return (
    <div>
      <ul>
        {props.participants.map(p => {
          return (
            <li>
              {p.login}: {p.role}
            </li>
          );
        })}
      </ul>
    </div>
  );
};
