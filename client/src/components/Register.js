import React, { useState } from "react"
import { TextField, Button } from "@material-ui/core"
import Alert from '@material-ui/lab/Alert';
import useInputState from "../hooks/useInputState"
import axios from "axios"
import { Navbar, Error } from "./";


function Register(props) {
    const [firstName, handleFirstNameChange] = useInputState("")
    const [lastName, handleLastNameChange] = useInputState("")
    const [email, handleEmailChange] = useInputState("")
    const [password, handlePasswordChange] = useInputState("")
    const [isRegistered, setIsRegistered] = useState(false)
    const [message, setMessage] = useState("")

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            await axios.post("http://localhost:8080/user/register", {
                firstName,
                lastName,
                email,
                password
            })
            setMessage("Great, you are almost there!")
            setIsRegistered(true)
        } catch (err) {
            setMessage(err.response.data.message)
        }
    }

    return (
        <div className="register">
            <Navbar history={props.history} />
            {
                !isRegistered ?
                    <>
                        <h1>Welcome to SampleAuth</h1>
                        <form
                            onSubmit={handleSubmit}
                            noValidate
                            autoComplete="off"
                            style={{
                                display: 'flex',
                                flexDirection: 'column'
                            }}
                        >
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
                            <TextField
                                label="Password"
                                value={password}
                                onChange={handlePasswordChange}
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
                        {
                            message !== "" &&
                            <Error message={message} setMessage={setMessage} />
                        }
                    </>
                    :
                    (
                        <div className="content">
                            <Alert
                                className="alert"
                                variant="filled"
                                severity="success"
                            >
                                {message}
                            </Alert>
                            <h3>Confirm your email address</h3>
                            <p>
                                We have sent an email with a confirmation link to your email address.
                                In order to complete the sign-up process, please click the confirmation link.
                            </p>
                        </div>
                    )
            }
        </div >
    )
}

export default Register