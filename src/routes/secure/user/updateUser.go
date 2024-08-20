package user

import (
	"controller/user"
	entities "domain/entity"
	"net/http"
	"util"
)

const (
	updateUserPattern string = "/user"
	updateUserMethod string = http.MethodPut
)

type UpdateUserRoute struct {
	Controller user.UpdateUserController
}

func (route UpdateUserRoute) Pattern() string {
	return updateUserPattern
}

func (route UpdateUserRoute) Method() string {
	return updateUserPattern
}

func (route UpdateUserRoute) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body entities.UserWithId
	err1 := util.GetRequestBody(r, &body)
	user, err2 := route.Controller.UpdateUser(&body)

	var err error
	if err1 == nil {
		err = err2
	}

	util.GetHttpResponse(w, r, user, err, true)	
}

