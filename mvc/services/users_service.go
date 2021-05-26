package services

import "github.com/GiovanniCavallari/golang-microservices/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
