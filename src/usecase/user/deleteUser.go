package user

import (
	"domain/entity"
	"repository/user"
)

type DeleteUserUsecase interface {
	DeleteUser(id int) (*domain.UserWithId, error)
}

type DeleteUserUsecaseImpl struct {
	Repository user.UserRepository
}

func (usecase DeleteUserUsecaseImpl) DeleteUser(id int) (*domain.UserWithId, error) {
	return usecase.Repository.DeleteById(id)
}
