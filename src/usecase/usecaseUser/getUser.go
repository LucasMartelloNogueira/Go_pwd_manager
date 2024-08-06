package usecaseUser

import (
	"domain"
	"userController"
)

func GetUser(id int) (domain.User, error) {
	return userController.GetUser(id)
}