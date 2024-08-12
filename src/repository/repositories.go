package repository

import (
	"repository/user"
)

var (
	UserRepository user.UserRepository = user.UserRepositoryCsv{}
)
