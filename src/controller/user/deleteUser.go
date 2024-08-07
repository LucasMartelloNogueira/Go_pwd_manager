package user

import (
	"domain"
	usecase "usecase/user"
)

func DeleteUser(id int) (domain.User, error) {
	return usecase.DeleteUser(id)
}