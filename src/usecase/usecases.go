package usecase

import (
	"usecase/auth"
	"usecase/user"
	repositories "repository"
)

var (
	LoginUsecase auth.LoginUsecase = auth.LoginUsecaseImpl{Repository: repositories.UserRepository}
	RegisterUsecase auth.RegisterUsecase = auth.RegisterUsecaseImpl{Repository: repositories.UserRepository}

	CreateUserUsecase user.CreateUserUsecase = user.CreateUserUsecaseImpl{Repository: repositories.UserRepository}
	DeleteUserUsecase user.DeleteUserUsecase = user.DeleteUserUsecaseImpl{Repository: repositories.UserRepository}
	GetUserUsecase user.GetUserUsecase = user.GetUserUsecaseImpl{Repository: repositories.UserRepository}
	UpdateUserUsecase user.UpdateUserUsecase = user.UpdateUserUsecaseImpl{Repository: repositories.UserRepository}
)