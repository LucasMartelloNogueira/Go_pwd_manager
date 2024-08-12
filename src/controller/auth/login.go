package auth

import (
	dtos "domain/dto"
	usecase "usecase/auth"
)

func Login(userCredentials dtos.UserCredentials) (*dtos.LoginResponse, error) {
	return usecase.Login(userCredentials)
}
