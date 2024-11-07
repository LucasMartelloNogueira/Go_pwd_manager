package user

import (
	"domain/entity"
)

type UserRepository interface {
	FindById(id int) (*domain.User, error)
	FindByColumn(column string, value string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	DeleteById(id int) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
}
