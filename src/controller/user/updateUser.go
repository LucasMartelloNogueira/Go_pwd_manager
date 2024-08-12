package user

import (
	"domain/entity"
	usecase "usecase/user"
)

func UpdateUser(user *domain.UserWithId) (*domain.UserWithId, error) {
	return usecase.UpdateUser(user)
}
