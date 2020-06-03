import React, { useContext } from "react"
import { UserProfileContext } from "../context/user"
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import axios from "axios";

function Navbar(props) {
    const { changeUserProfile } = useContext(UserProfileContext)

    const handleRedirect = (e) => {
        switch (e.target.innerText.toUpperCase()) {
            case 'REGISTER':
                props.history.push("register")
                break
            case 'LOGIN':
                props.history.push('login')
                break
            case 'PROFILE':
                props.history.push('profile')
                break
            default:
                props.history.push('/')
        }
    }

    const handleLogout = async () => {
        try {
            // expire cookie
            await axios.post(
                `http://localhost:8081/user/logout`,
                null,
                { withCredentials: true }
            )
            
            // remove access_token from localStorage
            localStorage.removeItem('access_token')
            props.history.push('/')

            // remove from user from context
            changeUserProfile(null) 
        } catch (err) {
            console.log(err.response)
        }
    }

    const path = props.history.location.pathname

    return (
        <AppBar position="static" id="navbar">
            <Toolbar>
                <Typography id="app-icon" variant="h6" onClick={handleRedirect}>
                    SampleAuth
                </Typography>
            </Toolbar>
            <Toolbar>
                {
                    path !== '/profile' &&
                    <>
                        <Button color="inherit" onClick={handleRedirect}>Login</Button>
                        <Button color="inherit" onClick={handleRedirect}>Profile</Button>
                    </>
                }
                {
                    path === '/profile' &&
                    <Button color="inherit" onClick={handleLogout}>Logout</Button>
                }
            </Toolbar>
        </AppBar>
    )
}

export default Navbar