package user

import (
	"controller/user"
	entities "domain/entity"
	"net/http"
	"util"
)

const (
	createUserPattern string = "/user"
	createUserMethod string = http.MethodPost 
)

type CreateUserRoute struct {
	Controller user.CreateUserController
}

func (route CreateUserRoute) Pattern() string {
	return createUserPattern
}

func (route CreateUserRoute) Method() string {
	return createUserMethod
}

func (route CreateUserRoute) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body entities.User
	err1 := util.GetRequestBody(r, &body)
	user, err2 := route.Controller.CreateUser(&body)

	var err error
	if err1 == nil {
		err = err2
	}

	util.GetHttpResponse(w, r, user, err, true)	
}

