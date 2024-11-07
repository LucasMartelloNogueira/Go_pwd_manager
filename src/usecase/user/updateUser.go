package user

import (
	entities "domain/entity"
	types "domain/types"
)

type UpdateUserUsecase interface {
	UpdateUser(id int, user *entities.User) (*entities.User, error)
}

type UpdateUserUsecaseImpl struct {
	Repository types.Repository[entities.User, entities.NewUser]
}

func (usecase UpdateUserUsecaseImpl) UpdateUser(id int, user *entities.User) (*entities.User, error) {
	return usecase.Repository.Update(id, user)
}
