package user

import (
	"domain"
	repositories "repository"
)

func GetUser(id int) (domain.User, error) {
	return repositories.UserRepository.FindById(id)
}