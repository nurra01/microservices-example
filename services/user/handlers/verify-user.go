package handlers

import (
	"context"
	"fmt"
	"net/http"
	"services/user/kafka"
	"services/user/models"
	"services/user/redis"
	"services/user/utils"

	"github.com/gorilla/mux"
)

// Verify handles GET request to verify a new user
func (h *RegisterUserHandler) Verify(rw http.ResponseWriter, req *http.Request) {
	verifyID := mux.Vars(req)["id"]
	h.log.Infof("verify user with id: %s\n", verifyID)

	// fetch user from redis using verifyID as a key
	user, err := redis.GetUser(verifyID)
	if err != nil || user == nil {
		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&models.GenericError{Message: "user not found, verification time expired"}, rw)
		h.log.Error(err)
		return
	}

	// all good and user is verified now
	user.Verified = true

	// convert user object to byte stream
	valStr, err := utils.FromObjectToByte(user)
	if err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		utils.ToJSON(&models.GenericError{Message: err.Error()}, rw)
		h.log.Fatal("failed to deserialize user to byte stream. %v", err.Error())
	}

	// write message to the kafka
	err = kafka.PushVerUser(context.Background(), []byte(user.Email), valStr)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&models.GenericError{Message: "failed to process user registration"}, rw)
		h.log.Fatal("failed writing a message to the kafka ", err)
	}

	rw.WriteHeader(http.StatusOK)
	utils.ToJSON(&models.GenericResponse{Response: fmt.Sprintf("user with email '%s' successfully verified", user.Email)}, rw)
}
