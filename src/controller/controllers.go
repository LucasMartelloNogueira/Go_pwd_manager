package controller

import (
	"controller/auth"
	"controller/user"
	usecases "usecase"
)

var (
	LoginController    auth.LoginController    = auth.LoginController{Usecase: usecases.LoginUsecase}
	RegisterController auth.RegisterController = auth.RegisterController{Usecase: usecases.RegisterUsecase}

	CreateUserController user.CreateUserController = user.CreateUserController{Usecase: usecases.CreateUserUsecase}
	DeleteUserController user.DeleteUserController = user.DeleteUserController{Usecase: usecases.DeleteUserUsecase}
	GetUserController    user.GetUserController    = user.GetUserController{Usecase: usecases.GetUserUsecase}
	UpdateUserController user.UpdateUserController = user.UpdateUserController{Usecase: usecases.UpdateUserUsecase}
)
