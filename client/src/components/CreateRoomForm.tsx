import * as React from "react";

import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

import "./Form.css";

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

export const CreateRoomForm = () => {
  const [roomName, setRoomName] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [login, setLogin] = React.useState("");

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    let xhr = new XMLHttpRequest();
    xhr.addEventListener("load", () => {
      console.log(xhr.responseText);
      onCreate();
    });
    xhr.open("POST", "http://localhost:8080/room");
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify({ name: roomName, password, login }));
  };

  const onCreate = () => {
    console.log("onCreate called");
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
