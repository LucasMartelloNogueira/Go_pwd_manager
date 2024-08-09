package routes

import (
	"domain"
	healthCheck "routes/public/helloWorld"
	"routes/public/auth"
	"routes/secure/user" 
)

var Routes []domain.Route = []domain.Route{
	healthCheck.HelloWorld,

	auth.Register,
	auth.Login,
	
	user.CreateUser,
	user.GetUser,
	user.UpdateUser,
	user.DeleteUser,
}

