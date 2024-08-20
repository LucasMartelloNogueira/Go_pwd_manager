package auth

import (
	"controller/auth"
	dtos "domain/dto"
	"net/http"
	"util"
)

const (
	loginPattern = "/auth/login"
	loginMethod = http.MethodPost
)

type LoginRoute struct {
	Controller auth.LoginController
}

func (route LoginRoute) Pattern() string {
	return loginPattern
}

func (route LoginRoute) Method() string {
	return loginMethod
}

func (route LoginRoute) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var userCredentials dtos.UserCredentials
	err1 := util.GetRequestBody(r, &userCredentials)
	loginInfo, err2 := route.Controller.Login(userCredentials)

	var err error
	if err1 == nil {
		err = err2
	}

	util.GetHttpResponse(w, r, loginInfo, err, false)
}
