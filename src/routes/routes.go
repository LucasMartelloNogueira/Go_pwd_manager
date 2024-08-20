package routes

import (
	types "domain/types"
	"routes/public/auth"
	"routes/public/helloWorld"

	"routes/secure/user"
	controllers "controller"
)

var (
	HelloWorldRoute types.Route = helloWorld.HelloWorldRoute{}

	LoginRoute types.Route = auth.LoginRoute{Controller: controllers.LoginController}
	RegisterRoute types.Route = auth.RegisterRoute{Controller: controllers.RegisterController}

	CreateUserRoute types.Route = user.CreateUserRoute{Controller: controllers.CreateUserController}
	DeleteUserRoute types.Route = user.DeleteUserRoute{Controller: controllers.DeleteUserController}
	GetUserRoute types.Route = user.GetUserRoute{Controller: controllers.GetUserController}
	UpdateUserRoute types.Route = user.UpdateUserRoute{Controller: controllers.UpdateUserController}
)

var Routes []types.Route = []types.Route{
	// healthcheck
	HelloWorldRoute,

	// auth
	LoginRoute,
	RegisterRoute,

	// user
	CreateUserRoute,
	DeleteUserRoute,
	GetUserRoute,
	UpdateUserRoute,
}
