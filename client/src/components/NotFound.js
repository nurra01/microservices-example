import React from "react"
import { Button } from "@material-ui/core"

function NotFound(props) {
    const handleRedirect = () => {
        props.history.push("/")
    }
    return (
        <div className="not-found">
            <div className="not-found-content">
                <h1 id="title">Ooops!</h1>
                <h1 id="message">404 - PAGE NOT FOUND</h1>
                <Button
                    type="submit"
                    variant="contained"
                    color="primary"
                    onClick={handleRedirect}
                >
                    Go to homepage
                </Button>
            </div>
        </div>
    )
}

export default NotFound