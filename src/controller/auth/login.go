package auth

import (
	dtos "domain/dto"
	"usecase/auth"
)

type LoginController interface {
	Login(userCredentials dtos.UserCredentials) (dtos.LoginResponse, error)
}

type LoginControllerImpl struct {
	Usecase auth.LoginUsecase
}


func (controller LoginControllerImpl) Login(userCredentials dtos.UserCredentials) (dtos.LoginResponse, error) {
	return controller.Usecase.Login(userCredentials)
}
