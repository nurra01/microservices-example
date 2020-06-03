package handlers

import (
	"net/http"
	"time"
)

// Logout resets cookie expire time to yesterday
func (h *AuthHandler) Logout(rw http.ResponseWriter, req *http.Request) {
	h.log.Println(req.Cookies())
	// access refresh token from session
	cookie, err := req.Cookie("refresh_token")
	if err != nil {
		RespondError(rw, "missing required cookie, please log in", http.StatusUnauthorized)
		return
	}

	cookie.Expires = time.Now().Add(-24 * time.Hour) // reset value to yesterday
	http.SetCookie(rw, cookie)                       // set new expired cookie

	RespondSuccessJSON(rw, nil, http.StatusNoContent)
}
