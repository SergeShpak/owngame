import * as React from "react";

import { CreateRoom } from "../pages/CreateRoom";
import * as Paths from "../paths";
import { Room } from "../pages/Room";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

export const Main: React.StatelessComponent<{}> = props => {
  return (
    <MainWrapper>
      <Router>
        <Switch>
          <Route exact path={"/"} component={CreateRoom} />
          <Route
            path={Paths.ROOM}
            render={props => <Room token={props.location.state.token} />}
          />
        </Switch>
      </Router>
    </MainWrapper>
  );
};

export const MainWrapper: React.StatelessComponent = props => (
  <main className="hero section">
    <div className="main">{props.children}</div>
  </main>
);
