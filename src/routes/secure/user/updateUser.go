package user

import (
	"net/http"

	controllers "controller"
	types "domain/types"
)

var UpdateUserRoute types.Route = types.Route{
	Pattern:    "/user",
	Method:     http.MethodPut,
	Controller: controllers.UpdateUserController,
}
