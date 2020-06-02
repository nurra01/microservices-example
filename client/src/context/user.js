import React, { useState } from "react"

export const UserProfileContext = React.createContext({})

export function UserProfileProvider(props) {
    const [userProfile, setUserProfile] = useState(null);

    const changeUserProfile = (user) => setUserProfile(user);

    return (
        <UserProfileContext.Provider value={{ userProfile, changeUserProfile }}>
            {props.children}
        </UserProfileContext.Provider>
    );
}