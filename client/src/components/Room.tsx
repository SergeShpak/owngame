import * as React from "react";

interface RoomProps {
  participants: string[];
}

export const Room: React.StatelessComponent<RoomProps> = props => {
  return (
    <div>
      <ul>
        {props.participants.map(p => {
          return <li>{p}</li>;
        })}
      </ul>
    </div>
  );
};
