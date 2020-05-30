import React from "react"
import { Button } from "@material-ui/core"

function Home(props) {
    const handleRedirect = (e) => {
        if (e.target.innerText === "REGISTER") {
            props.history.push("register")
        } else {
            props.history.push("login")
        }
    }
    return (
        <div className="home">
            <h1>Welcome to microservices example app</h1>
            <div className="home-btns">
                <Button
                    variant="contained"
                    color="primary"
                    onClick={handleRedirect}>
                    Register
                </Button>
                <Button
                    variant="contained"
                    color="primary"
                    onClick={handleRedirect}>
                    Login
                </Button>
            </div>
        </div>
    )
}

export default Home