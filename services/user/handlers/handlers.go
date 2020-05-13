package handlers

import (
	"github.com/sirupsen/logrus"
)

// RegisterUserHandler for registering users
type RegisterUserHandler struct {
	log *logrus.Logger
}

// NewUserHandler returns a new handler with the given dependencies
func NewUserHandler(log *logrus.Logger) *RegisterUserHandler {
	return &RegisterUserHandler{
		log,
	}
}
