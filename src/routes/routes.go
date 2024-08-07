package routes

import (
	"domain"
	"routes/greet"
	"routes/user" 
)

var Routes []domain.Route = []domain.Route{
	greet.Greet,
	user.CreateUser,
	user.GetUser,
	user.UpdateUser,
	user.DeleteUser,
}

