package auth

import (
	entities "domain/entity"
	"usecase/auth"
)

type RegisterController interface {
	Register(user *entities.User) (entities.UserWithId, error)
}

type RegisterControllerImpl struct {
	Usecase auth.RegisterUsecase
}

func (controller RegisterControllerImpl) Register(user *entities.User) (entities.UserWithId, error) {
	return controller.Usecase.Register(user)
}

