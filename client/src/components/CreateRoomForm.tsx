import * as React from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

import "./Form.css";
import { XHRRequest } from "../lib/xhr";

interface FormControlProps {
  value: string;
  setValue: React.Dispatch<React.SetStateAction<string>>;
  type?: string;
  placeholder?: string;
}

const FormControl: React.FunctionComponent<FormControlProps> = props => {
  return (
    <Form.Control
      value={props.value}
      onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
        props.setValue(e.target.value)
      }
      type={props.type}
      placeholder={props.placeholder}
    />
  );
};

export const CreateRoomForm = (props: { onCreate: () => void }) => {
  const [roomName, setRoomName] = React.useState("MyRoom");
  const [password, setPassword] = React.useState("P@ssw0rd");
  const [login, setLogin] = React.useState("BuPin");

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    XHRRequest.send({
      method: "POST",
      path: "room",
      body: {
        name: roomName,
        password,
        login
      }
    })
      .then(resp => {
        onCreate();
      })
      .catch(r => {
        throw r;
      });
  };

  const onCreate = () => {
    props.onCreate();
  };

  return (
    <div id="CreateRoomForm">
      <Form className="Form" onSubmit={onSubmit}>
        <Form.Group controlId="formLogin">
          <Form.Label>Login</Form.Label>
          <FormControl placeholder="Login" setValue={setLogin} value={login} />
        </Form.Group>
        <Form.Group controlId="formRoomName">
          <Form.Label>Room name</Form.Label>
          <FormControl
            placeholder="Room name"
            setValue={setRoomName}
            value={roomName}
          />
        </Form.Group>
        <Form.Group controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <FormControl
            type="password"
            placeholder="Password"
            setValue={setPassword}
            value={password}
          />
        </Form.Group>
        <Button variant="primary" type="submit">
          Join
        </Button>
      </Form>
    </div>
  );
};
