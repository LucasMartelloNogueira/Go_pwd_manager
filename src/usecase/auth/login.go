package auth

import (
	"CustomErrors"
	dtos "domain/dto"
	repositories "repository"
	"time"
	"util"
	"strconv"
)

func Login(userCredentials dtos.UserCredentials) (*dtos.LoginResponse, error){
	user, err := repositories.UserRepository.FindByColumn("email", userCredentials.Email)

	if err != nil {
		return nil, err
	}


	if strconv.Itoa(userCredentials.Password) != user.Password {
		return nil, CustomErrors.ErrBadRequest
	}

	tokenClaims := map[string]any {
		"userId": user.Id,
		"alg": util.Algorithm,
		"roles": []string{},
		"expiration": time.Now().Add(time.Hour),
	}

	token, err := util.JwtEncode(tokenClaims)
	if err != nil {
		return nil, CustomErrors.ErrInternalServerError
	}

	return &dtos.LoginResponse{
		UserId: user.Id,
		Token: token,
	}, nil
}