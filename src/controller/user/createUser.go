package user

import (
	"domain/entity"
	usecase "usecase/user"
)

type CreateUserController interface {
	CreateUser(user *domain.User) (*domain.UserWithId, error)
}

type CreateUserControllerImpl struct {
	Usecase usecase.CreateUserUsecase
}

func (controller CreateUserControllerImpl) CreateUser(user *domain.User) (*domain.UserWithId, error) {
	return controller.Usecase.CreateUser(user)
}

// func CreateUser(user *domain.User) (*domain.UserWithId, error) {
// 	return usecase.CreateUser(user)
// }
