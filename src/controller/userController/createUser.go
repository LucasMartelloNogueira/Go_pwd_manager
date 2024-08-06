package userController

import (
	"domain"
)


func CreateUser(createUserBody *domain.CreateUserBody) (domain.User, error){
	return userController.repository.Create(createUserBody)
}