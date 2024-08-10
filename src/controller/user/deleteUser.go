package user

import (
	"domain/entity"
	usecase "usecase/user"
)

func DeleteUser(id int) (*domain.UserWithId, error) {
	return usecase.DeleteUser(id)
}