package user

import (
	entities "domain/entity"
	types "domain/types"
)

type GetUserUsecase interface {
	GetUser(id int) (*entities.User, error)
}

type GetUserUsecaseImpl struct {
	Repository types.Repository[entities.User, entities.NewUser]
}

func (usecase GetUserUsecaseImpl) GetUser(id int) (*entities.User, error) {
	return usecase.Repository.FindById(id)
}
