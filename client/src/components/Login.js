import React, { useState, useEffect } from "react"
import { TextField, Button } from "@material-ui/core"
import Alert from '@material-ui/lab/Alert';
import useInputState from "../hooks/useInputState"
import axios from "axios"

function Login(props) {
    const [email, handleEmailChange] = useInputState("")
    const [password, handlePasswordChange] = useInputState("")
    const [isLoggedIn, setIsLoggedIn] = useState(false)
    const [message, setMessage] = useState("")

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            await axios.post("http://localhost:8081/user/login", {
                email,
                password
            })
            setIsLoggedIn(true)
        } catch (err) {
            setMessage(err.response.data.message)
        }
    }

    useEffect(() => {
        if (isLoggedIn) {
            props.history.push("/")
        }
    })

    return (
        <div className="register">
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
                {
                    message !== "" &&
                    <Alert
                        className="alert"
                        variant="filled"
                        severity={"error"}
                        onClose={() => {
                            setMessage("")
                        }}
                    >
                        {message}
                    </Alert>
                }
            </>
        </div >
    )
}

export default Login