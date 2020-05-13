package models

// GenericError defines error response for the request
type GenericError struct {
	Message string `json:"message"`
}

// GenericResponse defines successfull response for the request
type GenericResponse struct {
	Response string `json:"response"`
}
