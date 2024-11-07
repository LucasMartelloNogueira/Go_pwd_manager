package auth

import (
	"CustomErrors"
	entities "domain/entity"
	types "domain/types"
)

type RegisterUsecase interface {
	Register(newUser *entities.NewUser) (*entities.User, error)
}

type RegisterUsecaseImpl struct {
	Repository types.Repository[entities.User, entities.NewUser]
}

func (usecase RegisterUsecaseImpl) Register(newUser *entities.NewUser) (*entities.User, error) {
	user, err := usecase.Repository.Create(newUser)

	if err != nil {
		return nil, CustomErrors.ErrInternalServerError
	}

	return user, nil
}
