package user

import (
	"domain/entity"
	repositories "repository"
)

func UpdateUser(user *domain.UserWithId) (*domain.UserWithId, error) {
	return repositories.UserRepository.Update(user)
}