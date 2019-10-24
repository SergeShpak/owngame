import * as React from "react";

import { CreateRoom } from "../pages/CreateRoom";

export const Main: React.StatelessComponent<{}> = props => {
  return (
    <MainWrapper>
      <CreateRoom />
    </MainWrapper>
  );
};

export const MainWrapper: React.StatelessComponent = props => (
  <main className="hero section">
    <div className="main">{props.children}</div>
  </main>
);
