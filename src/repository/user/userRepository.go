package user

import (
	"domain/entity"
)

type UserRepository interface {
	FindById(id int) (*domain.UserWithId, error)
	FindByColumn(column string, value string) (*domain.UserWithId, error)
	Create(user *domain.User) (*domain.UserWithId, error)
	DeleteById(id int) (*domain.UserWithId, error)
	Update(user *domain.UserWithId) (*domain.UserWithId, error)
}
