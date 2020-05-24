package handlers

import "github.com/sirupsen/logrus"

// AuthHandler for authentication
type AuthHandler struct {
	log *logrus.Logger
}

// NewAuthHandler returns a new handler with the given dependencies
func NewAuthHandler(log *logrus.Logger) *AuthHandler {
	return &AuthHandler{log}
}
