package user

import (
	"domain"
	repositories "repository"
)

func UpdateUser(user *domain.User) (domain.User, error) {
	return repositories.UserRepository.Update(*user)
}