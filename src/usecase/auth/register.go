package auth

import (
	"CustomErrors"
	entities "domain/entity"
	"repository/user"
)

type RegisterUsecase interface {
	Register(user *entities.User) (entities.UserWithId, error) 
}

type RegisterUsecaseImpl struct {
	Repository user.UserRepository
}

func (usecase RegisterUsecaseImpl) Register(user *entities.User) (entities.UserWithId, error) {
	newUser, err := usecase.Repository.Create(user)

	if err != nil {
		return entities.UserWithIdBuilder(), CustomErrors.ErrInternalServerError
	}

	return entities.UserWithIdBuilder(
				entities.WithId(newUser.Id), 
				entities.WithName(newUser.Name), 
				entities.WithEmail(newUser.Email), 
				entities.WithPassWord(newUser.Password),
			), nil
}
