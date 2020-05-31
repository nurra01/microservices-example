package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"services/auth/db"
	"services/auth/models"
	"services/auth/utils"
)

// KeyAccessTokenReq is context key for access token request
type KeyAccessTokenReq struct{}

// MiddlewareValidateAccessToken for validating access token request
func (h *AuthHandler) MiddlewareValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// get access token from header
		bearerToken, err := utils.ExtractAccessToken(req)
		if err != nil {
			RespondError(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		// validate access token
		accessTokenClaims, err := utils.TokenValid(bearerToken)
		// if token is expired check refresh token
		// if token is valid return existing valid token
		if err != nil {
			if err.Error() != "Token is expired" {
				RespondError(rw, err.Error(), http.StatusUnauthorized)
				return
			}
			// access refresh token from session
			cookie, err := req.Cookie("refresh_token")
			if err != nil {
				RespondError(rw, "missing required cookie, please log in", http.StatusUnauthorized)
				return
			}

			// validate refresh token to be sure it's still valid
			refreshTokenClaims, err := utils.TokenValid(cookie.Value)
			if err != nil {
				RespondError(rw, err.Error(), http.StatusUnauthorized)
				return
			}

			// validate both user ids from claims to be sure accessToken wasn't changed
			if accessTokenClaims["userID"] != refreshTokenClaims["userID"] {
				RespondError(rw, "Invalid tokens, please log in", http.StatusUnauthorized)
				return
			}

			// add a a login request to the context
			ctx := context.WithValue(req.Context(), KeyAccessTokenReq{}, accessTokenClaims["userID"])
			req = req.WithContext(ctx)

			// if all good return next handler or middleware
			next.ServeHTTP(rw, req)
		} else {
			accessTokenResp := &models.AccessTokenResp{
				Token: bearerToken,
			}
			// respond with JSON body
			RespondSuccessJSON(rw, accessTokenResp, http.StatusOK)
		}
	})
}

// AccessToken handles a request to get a new access token
func (h *AuthHandler) AccessToken(rw http.ResponseWriter, req *http.Request) {
	// get userID from request context
	userID := req.Context().Value(KeyAccessTokenReq{}).(string)

	// get user from DB by id
	user, err := db.GetUserByID(userID)
	if err != nil {
		h.log.Errorf("failed to get user from DB. Error: %s", err.Error())
		if err == sql.ErrNoRows {
			RespondError(rw, "invalid email/password", http.StatusBadRequest)
		} else {
			RespondError(rw, "failed to process request", http.StatusInternalServerError)
		}
		return
	}

	// Generate a JWT access token
	token, err := utils.GenerateToken(user)
	if err != nil {
		RespondError(rw, "failed to process request", http.StatusInternalServerError)
		return
	}

	accessTokenResp := &models.AccessTokenResp{
		Token: token,
	}

	// respond with JSON body
	RespondSuccessJSON(rw, accessTokenResp, http.StatusOK)
}
