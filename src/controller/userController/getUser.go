package userController

import (
	"domain"
)


func GetUser(id int) (domain.User, error){
	return userController.repository.FindById(id);
}