package models

import (
	"errors"
	"regexp"
)

// GenericError defines error response for the request
type GenericError struct {
	Message string `json:"message"`
}

// GenericResponse defines successfull response for the request
type GenericResponse struct {
	Response string `json:"response"`
}

// LoginReq defines request structure for Login handler
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResp defines response stcture for Login request
type LoginResp struct {
	*User
	Token string `json:"access_token"`
}

// AccessTokenResp defines response structure for AccessToken request
type AccessTokenResp struct {
	Token string `json:"access_token"`
}

// Validate fields to be correct
func (l *LoginReq) Validate() error {
	if l.Email == "" || l.Password == "" {
		return errors.New("missing required body fields")
	}

	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !Re.MatchString(l.Email) {
		return errors.New("email is invalid")
	}
	return nil
}
