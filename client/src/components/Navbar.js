import React from "react"
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';

function Navbar(props) {
    const handleRedirect = (e) => {
        switch (e.target.innerText.toUpperCase()) {
            case 'REGISTER':
                props.history.push("register")
                break
            case 'LOGIN':
                props.history.push('login')
                break
            default:
                props.history.push('profile')
        }
    }

    return (
        <AppBar position="static" id="navbar">
            <Toolbar>
                <Typography variant="h6">
                    SampleAuth
            </Typography>
            </Toolbar>
            <Toolbar>
                <Button color="inherit" onClick={handleRedirect}>Login</Button>
                <Button color="inherit" onClick={handleRedirect}>Profile</Button>
            </Toolbar>
        </AppBar>
    )
}

export default Navbar