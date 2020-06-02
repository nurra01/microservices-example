import React, { useState, useEffect, useContext } from "react"
import { UserProfileContext } from "../context/user"
import axios from "axios"
import jwt from "jsonwebtoken"
import { setAccessToken } from "../utils/utils"
import { Navbar, Error } from "./"
import CircularProgress from '@material-ui/core/CircularProgress';

function Profile(props) {
    const { userProfile, changeUserProfile } = useContext(UserProfileContext)
    const [message, setMessage] = useState("")
    const token = localStorage.getItem('access_token')
    const decodedToken = jwt.decode(token)

    const options = {
        withCredentials: true,
        headers: {
            Authorization: `Bearer ${token}`
        }
    }

    const refreshToken = async () => {
        try {
            const resp = await axios.get(`http://localhost:8081/auth/token`, options)
            setAccessToken(resp.data['access_token'])
        } catch (err) {
            setMessage(err.response.data['message'])
            setTimeout(() => {
                props.history.push('login')
            }, 2000)
        }
    }

    const fetchProfile = async () => {
        try {
            const resp = await axios.get(`http://localhost:8081/users/${decodedToken.userID}/profile`, options)
            changeUserProfile(resp.data)
        } catch (err) {
            if (err.response.data['message'] === 'Token is expired') {
                await refreshToken()
            }
        }
    }

    useEffect(() => {
        if (decodedToken !== null) {
            fetchProfile()
        } else {
            setMessage('missing required token, please log in')
            setTimeout(() => {
                props.history.push('login')
            }, 2000)
        }
    }, [])

    return (
        <div className="profile">
            <Navbar history={props.history} />
            {userProfile !== null ?
                (
                    <div className="content">
                        <h1>User profile:</h1>
                        <h3>First name: {userProfile.firstName}</h3>
                        <h3>Last name: {userProfile.lastName}</h3>
                        <h3>Email: {userProfile.email}</h3>
                    </div>
                )
                :
                (
                    <div className="content">
                        <p>Please wait, loading your profile ...</p>
                        <CircularProgress size="4rem" />
                    </div>
                )
            }
            {
                message !== "" &&
                <Error message={message} setMessage={setMessage} />
            }
        </div>
    )
}

export default Profile