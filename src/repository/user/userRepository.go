package user

import (
	"domain"
)

type UserRepository interface {
	FindById(id int) (domain.User, error)
	Create(user *domain.CreateUserBody) (domain.User, error)
	DeleteById(id int) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}
