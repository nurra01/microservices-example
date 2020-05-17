import React from "react"
import { Button } from "@material-ui/core"

function Home(props) {
    const handleRedirect = () => {
        props.history.push("register")
    }
    return (
        <div className="home">
            <h1>Welcome to microservices example app</h1>
            <Button
                variant="contained"
                color="primary"
                onClick={handleRedirect}>
                Register
            </Button>
        </div>
    )
}

export default Home