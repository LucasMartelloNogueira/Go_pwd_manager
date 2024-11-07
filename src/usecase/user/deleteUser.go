package user

import (
	entities "domain/entity"
	types "domain/types"
)

type DeleteUserUsecase interface {
	DeleteUser(id int) (*entities.User, error)
}

type DeleteUserUsecaseImpl struct {
	Repository types.Repository[entities.User, entities.NewUser]
}

func (usecase DeleteUserUsecaseImpl) DeleteUser(id int) (*entities.User, error) {
	return usecase.Repository.DeleteById(id)
}
