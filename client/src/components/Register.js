import React, { useState } from "react"
import { TextField, Button } from "@material-ui/core"
import Alert from '@material-ui/lab/Alert';
import useInputState from "../hooks/useInputState"
import axios from "axios"


function Register() {
    const [firstName, handleFirstNameChange] = useInputState("")
    const [lastName, handleLastNameChange] = useInputState("")
    const [email, handleEmailChange] = useInputState("")
    const [password, handlePasswordChange] = useInputState("")
    const [isRegistered, setIsRegistered] = useState(null)
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
            setIsRegistered(false)
        }
    }

    return (
        <div className="register">
            {
                !isRegistered ?
                    <>
                        <h1>Register</h1>
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
                            isRegistered == false &&
                            <Alert
                                className="alert"
                                variant="filled"
                                severity={"error"}
                                onClose={() => {
                                    setIsRegistered(null)
                                }}
                            >
                                {message}
                            </Alert>
                        }
                    </>
                    :
                    <>
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
                    </>
            }
        </div >
    )
}

export default Register