package auth

import (
	"controller/auth"
	entities "domain/entity"
	"net/http"
	"util"
)

const (
	registerPattern = "/auth/register"
	registerMethod = http.MethodPost
)

type RegisterRoute struct {
	Controller auth.RegisterController
}

func (route RegisterRoute) Pattern() string {
	return registerPattern
}

func (route RegisterRoute) Method() string {
	return registerMethod
}

func (route RegisterRoute) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	util.GetRequestBody(r, &user)

	userWithId, err := route.Controller.Register(&user)
	util.GetHttpResponse(w, r, userWithId, err, false)
}

