package user

import (
	"domain"
	usecase "usecase/user"
)


func GetUser(id int) (domain.User, error){
	return usecase.GetUser(id)
}