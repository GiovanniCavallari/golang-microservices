package domain

import (
	"fmt"
	"net/http"

	"github.com/GiovanniCavallari/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, Firstname: "John", Lastname: "Doe", Email: "email@email.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User with id %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
