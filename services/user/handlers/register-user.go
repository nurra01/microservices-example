package handlers

import (
	"context"
	"net/http"

	"services/user/kafka"
	"services/user/models"
	"services/user/utils"

	"github.com/google/uuid"
)

// KeyUser is context key for user object
type KeyUser struct{}

// MiddlewareValidateRegisterUser validates all fields to be passed correctly
func (h *RegisterUserHandler) MiddlewareValidateRegisterUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// add header to make content JSON
		rw.Header().Add("Content-Type", "application/json")

		// user object
		usr := &models.RegisterUser{}

		// read req body and deserialize it to the
		err := utils.FromJSON(usr, req.Body)
		if err != nil {
			h.log.Println("failed deserializing user body", err)
			rw.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&models.GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the user
		err = usr.Validate()
		if err != nil {
			h.log.Error("failed validating user body")
			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			utils.ToJSON(&models.GenericError{Message: err.Error()}, rw)
			return
		}

		// hash user password
		usr.Password, err = utils.HashPassword(usr.Password)
		if err != nil {
			h.log.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&models.GenericError{Message: "failed to process passed password"}, rw)
			return
		}

		// add a user to the context
		ctx := context.WithValue(req.Context(), KeyUser{}, usr)
		req = req.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, req)
	})
}

// Register handles POST requests to add a new user
func (h *RegisterUserHandler) Register(rw http.ResponseWriter, req *http.Request) {
	// get user from req context after middleware validation
	usr := req.Context().Value(KeyUser{}).(*models.RegisterUser)

	// assign uuid
	usr.ID = uuid.New().String()

	// convert user object to byte stream
	valStr, err := utils.FromObjectToByte(usr)
	if err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		utils.ToJSON(&models.GenericError{Message: err.Error()}, rw)
		h.log.Fatal("failed to deserialize user to byte stream. %v", err.Error())
	}

	// write message to the kafka
	err = kafka.PushRegUser(context.Background(), []byte(usr.Email), valStr)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&models.GenericError{Message: "failed to process user registration"}, rw)
		h.log.Fatal("failed writing a message to the kafka ", err)
	}
}
