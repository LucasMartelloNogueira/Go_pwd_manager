package user

import (
	"domain"
	usecase "usecase/user"
)

func UpdateUser(user *domain.User) (domain.User, error) {
	return usecase.UpdateUser(user)
}