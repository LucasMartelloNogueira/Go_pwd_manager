package user

import (
	controllers "controller"
	types "domain/types"
	"net/http"
)

var CreateUserRoute types.Route = types.Route{
	Pattern:    "/user",
	Method:     http.MethodPost,
	Controller: controllers.CreateUserController,
}
