import React from 'react';
import { TextField, Button } from "@material-ui/core"
import useInputState from "./hooks/useInputState"
import axios from "axios"
import './App.css';

function App() {
  const [firstName, handleFirstNameChange] = useInputState("")
  const [lastName, handleLastNameChange] = useInputState("")
  const [email, handleEmailChange] = useInputState("")

  const handleSubmit = (event) => {
    axios.post("http://localhost:8080/user/register", {
      firstName,
      lastName,
      email
    }, {
      headers: {
        'Access-Control-Allow-Origin': '*'
      }
    }).then(val => {
      console.log(val)
      alert(val.statusText)
    }).catch(err => {
      console.error(err.response)
      if (err.response.data) {
        alert(err.response.data.message)
      }
    })
    event.preventDefault();
  }

  return (
    <div
      className="App"
      style={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
      }}>
      <h1>Welcome to microservices with kafka example</h1>
      <form
        onSubmit={handleSubmit}
        noValidate
        autoComplete="off"
        style={{
          width: '40%',
          display: 'flex',
          flexDirection: 'column'
        }}>
        <TextField
          label="First name"
          value={firstName}
          onChange={handleFirstNameChange}
        />
        <TextField
          label="Last name"
          value={lastName}
          onChange={handleLastNameChange}
        />
        <TextField
          label="Email"
          value={email}
          onChange={handleEmailChange}
        />
        <Button
          type="submit"
          variant="contained"
          color="primary"
          style={{
            width: '100px',
            margin: '20px auto'
          }}>
          Sign up
        </Button>
      </form>
    </div >
  );
}

export default App;
