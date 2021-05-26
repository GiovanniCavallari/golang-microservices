package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GiovanniCavallari/golang-microservices/mvc/services"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userIdParam := request.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		// Return Bad Request to the client
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Param user_id must be a number"))
		return
	}

	user, err := services.GetUser(userId)

	if err != nil {
		// Handle the err and return to the client
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))
		return
	}

	// Return user to the client
	jsonValue, _ := json.Marshal(user)

	response.Write(jsonValue)
}
