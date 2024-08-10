package user

import (
	"domain/entity"
	repositories "repository"
)

func DeleteUser(id int) (*domain.UserWithId, error) {
	return repositories.UserRepository.DeleteById(id)
}