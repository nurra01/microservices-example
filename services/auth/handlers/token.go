package handlers

import (
	"net/http"
	"services/auth/utils"
)

// middleware for validating access token
func validateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// extract token value
		bearerToken, err := utils.ExtractAccessToken(req)
		if err != nil {
			RespondError(rw, err.Error(), http.StatusBadRequest)
			return
		}

		// validate access token
		err = utils.TokenValid(bearerToken)
		if err != nil {
			RespondError(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		// run next handler or middleware
		next.ServeHTTP(rw, req)
	})
}

// AccessToken handles a request for accessing a token
func (h *AuthHandler) AccessToken(rw http.ResponseWriter, req *http.Request) {

}
