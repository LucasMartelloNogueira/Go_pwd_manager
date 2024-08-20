package auth

import (
	"CustomErrors"
	dtos "domain/dto"
	"repository/user"
	"strconv"
	"time"
	"util"
)

type LoginUsecase interface {
	Login(userCredentials dtos.UserCredentials) (dtos.LoginResponse, error)
}

type LoginUsecaseImpl struct {
	Repository user.UserRepository
}

func (usecase LoginUsecaseImpl) Login(userCredentials dtos.UserCredentials) (dtos.LoginResponse, error) {
	user, err := usecase.Repository.FindByColumn("email", userCredentials.Email)

	if err != nil {
		return dtos.LoginResponseBuilder(), err
	}

	if strconv.Itoa(userCredentials.Password) != user.Password {
		return dtos.LoginResponseBuilder(), CustomErrors.ErrBadRequest
	}

	tokenClaims := map[string]any{
		"userId":     user.Id,
		"alg":        util.Algorithm,
		"roles":      []string{},
		"expiration": time.Now().Add(time.Hour),
	}

	token, err := util.JwtEncode(tokenClaims)
	if err != nil {
		return dtos.LoginResponseBuilder(), CustomErrors.ErrInternalServerError
	}

	return dtos.LoginResponseBuilder(dtos.WithUserId(user.Id), dtos.WithToken(token)), nil
}
