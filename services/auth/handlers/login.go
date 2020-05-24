package handlers

import (
	"context"
	"database/sql"
	"fmt"
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
				rw.WriteHeader(http.StatusBadRequest)
				utils.ToJSON(&models.GenericError{Message: "missing required body fields"}, rw)
			} else {
				rw.WriteHeader(http.StatusUnprocessableEntity)
				utils.ToJSON(&models.GenericError{Message: err.Error()}, rw)
			}
			return
		}

		// validate required body fields to be correct
		err = loginReq.Validate()
		if err != nil {
			h.log.Printf("failed login request validation. Error: %s", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&models.GenericError{Message: err.Error()}, rw)
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
			rw.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&models.GenericError{Message: "invalid email/password"}, rw)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			utils.ToJSON(&models.GenericError{Message: "failed to process request"}, rw)
		}
		return
	}
	// Verify passwords to match
	err = utils.VerifyPassword(user.Password, loginReq.Password)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&models.GenericError{Message: "invalid email/password"}, rw)
		return
	}
	fmt.Println(user, "logged user in")
}
