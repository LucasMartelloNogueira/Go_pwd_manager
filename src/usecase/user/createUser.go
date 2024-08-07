package user

import (
	"domain"
	repositories "repository"
	
)

func CreateUser(createUserBody *domain.CreateUserBody) (domain.User, error) {
	return repositories.UserRepository.Create(createUserBody)
}