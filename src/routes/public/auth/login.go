package auth

import (
	dtos "domain/dto"
	types "domain/types"
	controller "controller/auth"
	"net/http"
	"util"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var userCredentials dtos.UserCredentials;
	err1 := util.GetRequestBody(r, &userCredentials)
	loginInfo, err2 := controller.Login(userCredentials)

	var err error
	if err1 == nil {
		err = err2
	}

	util.GetHttpResponse(w, r, loginInfo, err, false)
}


var Login types.Route = types.Route{
	Pattern: "/auth/login",
	Method: http.MethodPost,
	Handler: loginHandler,
}