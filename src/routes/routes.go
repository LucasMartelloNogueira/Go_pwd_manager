package routes

import (
	types "domain/types"
	healthCheck "routes/public/helloWorld"
	"routes/public/auth"
	"routes/secure/user" 
)

var Routes []types.Route = []types.Route{
	healthCheck.HelloWorld,

	auth.Register,
	auth.Login,
	
	user.CreateUser,
	user.GetUser,
	user.UpdateUser,
	user.DeleteUser,
}

