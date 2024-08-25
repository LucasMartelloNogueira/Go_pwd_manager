package routes

import (
	types "domain/types"
	"routes/public/auth"
	"routes/public/helloWorld"
	"routes/secure/user"
)

var Routes []types.Route = []types.Route{
	// HealthCheck
	helloWorld.HelloWorldRoute,

	// Login
	auth.LoginRoute,
	auth.RegisterRoute,

	// user
	user.CreateUserRoute,
	user.DeleteUserRoute,
	user.GetUserRoute,
	user.UpdateUserRoute,
}
