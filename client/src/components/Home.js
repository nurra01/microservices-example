import React from "react"
import { Button } from "@material-ui/core"
import { Navbar } from "./"

function Home(props) {
    const handleRedirect = (e) => {
        props.history.push("register")
    }
    return (
        <div className="home">
            <Navbar history={props.history} />
            <div className="content">
                <h1>Welcome</h1>
                <p>This is simple authentication app which consumes microservices.</p>
                <p>To get started please sign up or log in to access your profile.</p>
            </div>

            <div className="home-btns">
                <Button
                    variant="contained"
                    color="primary"
                    onClick={handleRedirect}>
                    Sign Up
                </Button>
            </div>
        </div>
    )
}

export default Home