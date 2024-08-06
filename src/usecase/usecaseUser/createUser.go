package usecaseUser

import (
	"domain"
	"userController"
)

func CreateUser(createUserBody *domain.CreateUserBody) (domain.User, error) {
	return userController.CreateUser(createUserBody)
}