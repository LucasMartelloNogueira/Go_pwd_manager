package auth

import (
	controllers "controller"
	types "domain/types"
	"net/http"
)

var RegisterRoute types.Route = types.Route{
	Pattern:    "/auth/register",
	Method:     http.MethodPost,
	Controller: controllers.RegisterController,
}
