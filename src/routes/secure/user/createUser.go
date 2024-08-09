package user

import (
	"domain"
	"net/http"
	"util"
	controller "controller/user"
)


func createUserHandler(w http.ResponseWriter, r *http.Request){

	var body domain.CreateUserBody;
	util.GetRequestBody(r, &body)
	user, err := controller.CreateUser(&body);
	util.GetHttpResponse(w, r, user, err, true)
}

var CreateUser domain.Route = domain.Route{
	Pattern: "/user",
	Method: http.MethodPost,
	Handler: createUserHandler,
}