package user

import (
	controllers "controller"
	types "domain/types"
	"net/http"
)

var DeleteUserRoute types.Route = types.Route{
	Pattern:    "/user/{id}",
	Method:     http.MethodDelete,
	Controller: controllers.DeleteUserController,
}
