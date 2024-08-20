package user

import (
	"domain/entity"
	"usecase/user"
)

type GetUserController interface {
	GetUser(id int) (*domain.UserWithId, error)
}

type GetUserControllerImpl struct {
	Usecase user.GetUserUsecase
}

func (controller GetUserControllerImpl) GetUser(id int) (*domain.UserWithId, error) {
	return controller.Usecase.GetUser(id)
}
