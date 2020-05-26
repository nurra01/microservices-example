package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"services/auth/db"
	"services/auth/models"
	"services/auth/utils"
)

// KeyLoginReq is context key for login request object
type KeyLoginReq struct{}

// MiddlewareValidateLogin middleware with validation for Login handler
func (h *AuthHandler) MiddlewareValidateLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// add header to make content JSON
		rw.Header().Add("Content-Type", "application/json")

		loginReq := &models.LoginReq{}

		// read req body and deserialize it to the object
		err := utils.FromJSON(loginReq, req.Body)
		if err != nil {
			h.log.Printf("failed deserializing login request body. Error: %s", err.Error())
			if err.Error() == "EOF" {
				RespondError(rw, "missing required body fields", http.StatusBadRequest)
			} else {
				RespondError(rw, err.Error(), http.StatusUnprocessableEntity)
			}
			return
		}

		// validate required body fields to be correct
		err = loginReq.Validate()
		if err != nil {
			h.log.Printf("failed login request validation. Error: %s", err.Error())
			RespondError(rw, err.Error(), http.StatusBadRequest)
			return
		}

		// add a a login request to the context
		ctx := context.WithValue(req.Context(), KeyLoginReq{}, loginReq)
		req = req.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, req)
	})
}

// Login handler
func (h *AuthHandler) Login(rw http.ResponseWriter, req *http.Request) {
	// get login request from req context after middleware validation
	loginReq := req.Context().Value(KeyLoginReq{}).(*models.LoginReq)

	// get verified user from db
	user, err := db.GetUser(loginReq.Email)
	if err != nil {
		h.log.Errorf("failed to get user from DB. Error: %s", err.Error())
		if err == sql.ErrNoRows {
			RespondError(rw, "invalid email/password", http.StatusBadRequest)
		} else {
			RespondError(rw, "failed to process request", http.StatusInternalServerError)
		}
		return
	}
	// Verify passwords to match
	err = utils.VerifyPassword(user.Password, loginReq.Password)
	if err != nil {
		RespondError(rw, "invalid email/password", http.StatusBadRequest)
		return
	}

	// Generate a JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		RespondError(rw, "failed to process request", http.StatusInternalServerError)
		return
	}

	// Genereate a JWT refresh token
	refreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		RespondError(rw, "failed to process request", http.StatusInternalServerError)
		return
	}

	// response structure
	loginResp := &models.LoginResp{
		user,
		token,
		refreshToken,
	}
	RespondSuccessJSON(rw, loginResp, http.StatusOK)
}
