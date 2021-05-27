package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GiovanniCavallari/golang-microservices/mvc/services"
	"github.com/GiovanniCavallari/golang-microservices/mvc/utils"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userIdParam := request.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	response.Header().Set("Content-Type", "application/json")

	if err != nil {
		// Return Bad Request to the client
		apiErr := &utils.ApplicationError{
			Message:    "Param user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		// Handle the err and return to the client
		jsonValue, _ := json.Marshal(apiErr)

		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	// Return user to the client
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}
