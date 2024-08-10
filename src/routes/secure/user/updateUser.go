package user

import (
	entities "domain/entity"
	types "domain/types"
	"net/http"
	"util"
	controller "controller/user"

)

func updateUserHandler(w http.ResponseWriter, r *http.Request){
	var body entities.UserWithId;
	err1 := util.GetRequestBody(r, &body)
	user, err2 := controller.UpdateUser(&body);

	var err error
	if err1 == nil{
		err = err2
	}

	util.GetHttpResponse(w, r, user, err, true)
}

var UpdateUser types.Route = types.Route{
	Pattern: "/user",
	Method: http.MethodPut,
	Handler: updateUserHandler,
}