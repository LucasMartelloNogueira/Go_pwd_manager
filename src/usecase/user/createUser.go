package user

import (
	"domain/entity"
	repositories "repository"
)

func CreateUser(user *domain.User) (*domain.UserWithId, error) {
	return repositories.UserRepository.Create(user)
}
