package auth

import (
	"domain/entity"
	repository "repository"
)

func Register(user *domain.User) (*domain.UserWithId, error) {
	return repository.UserRepository.Create(user);
}