package handlers

import (
	"context"
	"net/http"
	"services/auth/db"
	"services/auth/models"
	"services/auth/utils"

	"github.com/gorilla/mux"
)

// KeyUserProfileReq is context key for user profile request
type KeyUserProfileReq struct{}

// MiddlewareValidateUserProfile middleware with validation for UserProfile handler
func (h *AuthHandler) MiddlewareValidateUserProfile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// get access token from header
		bearerToken, err := utils.ExtractAccessToken(req)
		if err != nil {
			RespondError(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		// validate access token
		accessTokenClaims, err := utils.TokenValid(bearerToken)
		if err != nil {
			if err.Error() == "Token is expired" {
				RespondError(rw, err.Error(), http.StatusUnauthorized)
			} else {
				RespondError(rw, "Invalid access token, please refresh it", http.StatusUnauthorized)
			}
			return
		}

		// add a a login request to the context
		ctx := context.WithValue(req.Context(), KeyUserProfileReq{}, accessTokenClaims["userID"])
		req = req.WithContext(ctx)

		// if all good return next handler or middleware
		next.ServeHTTP(rw, req)
	})
}

// UserProfile handles requests to get user information
func (h *AuthHandler) UserProfile(rw http.ResponseWriter, req *http.Request) {
	// get userID from request context
	userID := req.Context().Value(KeyUserProfileReq{}).(string)

	vars := mux.Vars(req)

	// compare user id passed as param with claim userID from token
	if userID != vars["id"] {
		RespondError(rw, "invalid user id", http.StatusUnauthorized)
		return
	}

	// get user from db by id
	user, err := db.GetUserByID(userID)
	if err != nil {
		h.log.Debug(err.Error())
		RespondError(rw, "Failed to proccess request", http.StatusInternalServerError)
		return
	}

	userProfile := &models.UserProfile{
		User:  user,
		Email: user.Email,
	}

	RespondSuccessJSON(rw, userProfile, http.StatusOK)
}
