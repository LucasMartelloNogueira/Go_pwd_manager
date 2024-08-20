package user

import (
	"domain/entity"
	"usecase/user"
)

type UpdateUserController interface {
	UpdateUser(user *domain.UserWithId) (*domain.UserWithId, error)
}

type UpdateUserControllerImpl struct {
	Usecase user.UpdateUserUsecase
}


func (controller UpdateUserControllerImpl) UpdateUser(user *domain.UserWithId) (*domain.UserWithId, error) {
	return controller.Usecase.UpdateUser(user)
}
