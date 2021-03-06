import React, { useState, useEffect } from "react"
import { CircularProgress, Button } from "@material-ui/core"
import axios from "axios"

function Verify(props) {
    const [isVerified, setIsVerified] = useState(null)
    const [message, setMessage] = useState("")
    const userID = props.match.params.userID


    const handleRedirect = () => {
        if (isVerified) {
            props.history.push("/login")
        } else {
            props.history.push("/register")
        }
    }

    const fetchAPI = async () => {
        try {
            const resp = await axios.get("http://localhost:8080/user/verify/" + userID)
            setMessage(resp.data.response)
            setIsVerified(true)
        } catch (err) {
            setMessage(err.response.data.message)
            setIsVerified(false)
        }
    }

    useEffect(() => {
        fetchAPI()
    }, [])

    return (
        <div className="verify">
            {isVerified == null ?
                (
                    <div className="content">
                        <h1>Verifing your account...</h1>
                        <CircularProgress size="4rem" />
                    </div>
                )
                :
                (
                    <div className="content">
                        <h1>{message}</h1>
                        <Button
                            type="submit"
                            variant="contained"
                            color="primary"
                            onClick={handleRedirect}
                        >
                            {isVerified ? "Login" : "Register"}
                        </Button>
                    </div>
                )
            }
        </div>
    )
}

export default Verify