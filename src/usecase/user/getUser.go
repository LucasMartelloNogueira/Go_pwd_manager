package user

import (
	"domain/entity"
	"repository/user"
	repositories "repository"
)

type GetUserUsecase interface {
	GetUser(id int) (*domain.UserWithId, error)
}

type GetUserUsecaseImpl struct {
	Repository user.UserRepository
}

func (usecase GetUserUsecaseImpl) GetUser(id int) (*domain.UserWithId, error) {
	return repositories.UserRepository.FindById(id)
}
