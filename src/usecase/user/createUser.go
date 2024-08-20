package user

import (
	"domain/entity"
	"repository/user"
)

type CreateUserUsecase interface {
	CreateUser(user *domain.User) (*domain.UserWithId, error)
}

type CreateUserUsecaseImpl struct {
	Repository user.UserRepository
}

func (usecase CreateUserUsecaseImpl) CreateUser(user *domain.User) (*domain.UserWithId, error) {
	return usecase.Repository.Create(user)
}
