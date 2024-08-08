package routes

import (
	"domain"
	"routes/greet"
	"routes/user" 
)

var Routes []domain.Route = []domain.Route{
	greet.HelloWorld,
	user.CreateUser,
	user.GetUser,
	user.UpdateUser,
	user.DeleteUser,
}

