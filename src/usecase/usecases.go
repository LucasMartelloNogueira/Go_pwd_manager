package usecase

import (
	repositories "repository"
	"usecase/auth"
	"usecase/user"
)

var (
	LoginUsecase    auth.LoginUsecase    = auth.LoginUsecaseImpl{Repository: repositories.CsvUserReposiory}
	RegisterUsecase auth.RegisterUsecase = auth.RegisterUsecaseImpl{Repository: repositories.CsvUserReposiory}

	CreateUserUsecase user.CreateUserUsecase = user.CreateUserUsecaseImpl{Repository: repositories.CsvUserReposiory}
	DeleteUserUsecase user.DeleteUserUsecase = user.DeleteUserUsecaseImpl{Repository: repositories.CsvUserReposiory}
	GetUserUsecase    user.GetUserUsecase    = user.GetUserUsecaseImpl{Repository: repositories.CsvUserReposiory}
	UpdateUserUsecase user.UpdateUserUsecase = user.UpdateUserUsecaseImpl{Repository: repositories.CsvUserReposiory}
)
