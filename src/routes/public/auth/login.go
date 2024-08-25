package auth

import (
	controllers "controller"
	types "domain/types"
	"net/http"
)

var LoginRoute types.Route = types.Route{
	Pattern:    "/auth/login",
	Method:     http.MethodPost,
	Controller: controllers.LoginController,
}
