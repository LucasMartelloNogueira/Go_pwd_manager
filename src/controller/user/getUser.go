package user

import (
	"domain/entity"
	usecase "usecase/user"
)


func GetUser(id int) (*domain.UserWithId, error){
	return usecase.GetUser(id)
}