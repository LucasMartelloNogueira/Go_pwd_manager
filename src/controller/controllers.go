package controller

import (
	"controller/auth"
	"controller/user"
	usecases "usecase"
)

var (
	LoginController auth.LoginController = auth.LoginControllerImpl{Usecase: usecases.LoginUsecase}
	RegisterController auth.RegisterController = auth.RegisterControllerImpl{Usecase: usecases.RegisterUsecase}

	CreateUserController user.CreateUserController = user.CreateUserControllerImpl{Usecase: usecases.CreateUserUsecase}
	DeleteUserController user.DeleteUserController = user.DeleteUserControllerImpl{Usecase: usecases.DeleteUserUsecase}
	GetUserController user.GetUserController = user.GetUserControllerImpl{Usecase: usecases.GetUserUsecase}
	UpdateUserController user.UpdateUserController = user.UpdateUserControllerImpl{Usecase: usecases.UpdateUserUsecase}
)