package services

import (
	"github.com/GiovanniCavallari/golang-microservices/mvc/domain"
	"github.com/GiovanniCavallari/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
