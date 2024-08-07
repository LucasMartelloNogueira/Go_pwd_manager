package user

import (
	"domain"
	repositories "repository"
)

func DeleteUser(id int) (domain.User, error) {
	return repositories.UserRepository.DeleteById(id)
}