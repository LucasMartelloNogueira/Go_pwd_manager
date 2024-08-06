package userController

import (
	"userRepository"
)

type UserController struct {
	repository userRepository.UserRepository 
}

var userController UserController = UserController{
	repository: userRepository.UserRepositoryCsv{},
}