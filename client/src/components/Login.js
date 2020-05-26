import React, { useState, useEffect } from "react"
import { TextField, Button } from "@material-ui/core"
import Alert from '@material-ui/lab/Alert';
import useInputState from "../hooks/useInputState"
import axios from "axios"

function Login(props) {
    const [email, handleEmailChange] = useInputState("")
    const [password, handlePasswordChange] = useInputState("")
    const [isLoggedIn, setIsLoggedIn] = useState(null)
    const [message, setMessage] = useState("")

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const resp = await axios.post("http://localhost:8081/user/login", {
                email,
                password
            })
            setIsLoggedIn(true)
        } catch (err) {
            setMessage(err.response.data.message)
            setIsLoggedIn(false)
        }
    }

    useEffect(() => {
        if (isLoggedIn) {
            props.history.push("/")
        }
    })

    return (
        <div className="register">
            {
                !isLoggedIn ?
                    <>
                        <h1>Log in</h1>
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

export default Login