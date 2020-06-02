import React, { useState, useEffect } from "react"
import { TextField, Button } from "@material-ui/core"
import useInputState from "../hooks/useInputState"
import axios from "axios"
import { setAccessToken } from "../utils/utils";
import { Navbar, Error } from "./"

function Login(props) {
    const [email, handleEmailChange] = useInputState("")
    const [password, handlePasswordChange] = useInputState("")
    const [isLoggedIn, setIsLoggedIn] = useState(false)
    const [message, setMessage] = useState("")

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const resp = await axios.post("http://localhost:8081/user/login",
                {
                    email,
                    password,
                },
                {
                    withCredentials: true
                }
            )
            // set access_token in session_storage
            setAccessToken(resp.data['access_token'])
            setIsLoggedIn(true)
        } catch (err) {
            setMessage(err.response.data.message)
        }
    }

    useEffect(() => {
        if (isLoggedIn) {
            props.history.push("/profile")
        }
    })

    return (
        <div className="register">
            <Navbar history={props.history} />
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
                        Log in
                    </Button>
                </form>
                {
                    message !== "" &&
                    <Error message={message} setMessage={setMessage} />
                }
            </>
        </div >
    )
}

export default Login