package user

import (
	"domain/entity"
	"usecase/user"
)

type DeleteUserController interface {
	DeleteUser(id int) (*domain.UserWithId, error)
}

type DeleteUserControllerImpl struct {
	Usecase user.DeleteUserUsecase
}

func (controller DeleteUserControllerImpl) DeleteUser(id int) (*domain.UserWithId, error) {
	return controller.Usecase.DeleteUser(id)
}
