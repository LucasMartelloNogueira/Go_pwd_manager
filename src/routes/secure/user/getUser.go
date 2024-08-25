package user

import (
	controllers "controller"
	types "domain/types"
	"net/http"
)

var GetUserRoute types.Route = types.Route{
	Pattern:    "/user/{id}",
	Method:     http.MethodGet,
	Controller: controllers.GetUserController,
}
