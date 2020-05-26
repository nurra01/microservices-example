package handlers

import (
	"net/http"
	"services/auth/models"
	"services/auth/utils"

	"github.com/sirupsen/logrus"
)

// AuthHandler for authentication
type AuthHandler struct {
	log *logrus.Logger
}

// NewAuthHandler returns a new handler with the given dependencies
func NewAuthHandler(log *logrus.Logger) *AuthHandler {
	return &AuthHandler{log}
}

// RespondError returns error message with status to the user
func RespondError(rw http.ResponseWriter, message string, status int) {
	rw.WriteHeader(status)
	utils.ToJSON(&models.GenericError{Message: message}, rw)
}

// RespondSuccessMessage returns successfull message with status to the user
func RespondSuccessMessage(rw http.ResponseWriter, message string, status int) {
	rw.WriteHeader(status)
	utils.ToJSON(&models.GenericResponse{Response: message}, rw)
}

// RespondSuccessJSON returns successfull JSON response body with status to the user
func RespondSuccessJSON(rw http.ResponseWriter, body interface{}, status int) {
	rw.WriteHeader(status)
	utils.ToJSON(body, rw)
}
