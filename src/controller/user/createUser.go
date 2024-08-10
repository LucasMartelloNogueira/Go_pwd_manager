package user

import (
	"domain/entity"
	usecase "usecase/user"
)


func CreateUser(user *domain.User) (*domain.UserWithId, error){
	return usecase.CreateUser(user)
}