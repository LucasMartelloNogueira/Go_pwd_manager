package routes

import (
	types "domain/types"
	"routes/public/auth"
	healthCheck "routes/public/helloWorld"
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
