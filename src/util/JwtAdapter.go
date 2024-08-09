package util

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"CustomErrors"
)

func jwtEncode(tokenClaims jwt.MapClaims) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	
	if err != nil {
		return "", CustomErrors.ErrInternalServerError
	}

	return tokenString, nil
}


func jwtDecode(tokenString string) (*jwt.Token, error) {
	decodedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, CustomErrors.ErrUnexpectedSigningMehthod
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	
	if err != nil {
		return nil, CustomErrors.ErrInternalServerError
	}

	return decodedToken, nil
}