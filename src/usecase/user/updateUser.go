package user

import (
	"domain/entity"
	"repository/user"
)

type UpdateUserUsecase interface {
	UpdateUser(user *domain.UserWithId) (*domain.UserWithId, error)
}

type UpdateUserUsecaseImpl struct {
	Repository user.UserRepository
}

func (usecase UpdateUserUsecaseImpl) UpdateUser(user *domain.UserWithId) (*domain.UserWithId, error) {
	return usecase.Repository.Update(user)
}
