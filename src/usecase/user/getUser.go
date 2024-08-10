package user

import (
	"domain/entity"
	repositories "repository"
)

func GetUser(id int) (*domain.UserWithId, error) {
	return repositories.UserRepository.FindById(id)
}