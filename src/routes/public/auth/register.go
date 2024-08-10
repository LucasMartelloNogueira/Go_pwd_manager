package auth

import (
	entities "domain/entity"
	types "domain/types"
	"net/http"
	"util"
	controller "controller/auth"
)


func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	util.GetRequestBody(r, &user)

	userWithId, err := controller.Register(&user);
	util.GetHttpResponse(w, r, userWithId, err, true)
}


var Register types.Route = types.Route{
	Pattern: "/auth/register",
	Method: http.MethodPost,
	Handler: registerHandler,
}