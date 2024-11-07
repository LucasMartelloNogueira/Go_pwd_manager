package user

import (
	entities "domain/entity"
	types "domain/types"
)

type CreateUserUsecase interface {
	CreateUser(newUser *entities.NewUser) (*entities.User, error)
}

type CreateUserUsecaseImpl struct {
	Repository types.Repository[entities.User, entities.NewUser]
}

func (usecase CreateUserUsecaseImpl) CreateUser(newUser *entities.NewUser) (*entities.User, error) {
	return usecase.Repository.Create(newUser)
}
