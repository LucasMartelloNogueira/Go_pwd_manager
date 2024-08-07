package user

import (
	"domain"
	usecase "usecase/user"
)


func CreateUser(createUserBody *domain.CreateUserBody) (domain.User, error){
	return usecase.CreateUser(createUserBody)
}