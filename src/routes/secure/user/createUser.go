package user

import (
	entities "domain/entity"
	types "domain/types"
	"net/http"
	"util"
	controller "controller/user"
)


func createUserHandler(w http.ResponseWriter, r *http.Request){
	var body entities.User;
	err1 := util.GetRequestBody(r, &body)
	user, err2 := controller.CreateUser(&body);

	var err error
	if err1 == nil {
		err = err2
	}

	util.GetHttpResponse(w, r, user, err, true)
}

var CreateUser types.Route = types.Route{
	Pattern: "/user",
	Method: http.MethodPost,
	Handler: createUserHandler,
}