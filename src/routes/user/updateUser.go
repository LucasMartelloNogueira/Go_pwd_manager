package user

import (
	"domain"
	"net/http"
	"util"
	controller "controller/user"

)

func updateUserHandler(w http.ResponseWriter, r *http.Request){
	var body domain.User;
	util.GetRequestBody(r, &body)
	user, err := controller.UpdateUser(&body);
	util.GetHttpResponse(w, r, user, err)
}

var UpdateUser domain.Route = domain.Route{
	Pattern: "/user",
	Method: http.MethodPut,
	Handler: updateUserHandler,
}