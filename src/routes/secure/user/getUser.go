package user

import (
	"controller/user"
	"net/http"
	"strconv"
	"util"
)

const (
	getUserPattern string = "/user/{id}"
	getUserMethod = http.MethodGet
)

type GetUserRoute struct {
	Controller user.GetUserController
}

func (route GetUserRoute) Pattern() string {
	return getUserPattern
}

func (route GetUserRoute) Method() string {
	return getUserMethod
}

func (route GetUserRoute) HandleRequest(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := route.Controller.GetUser(id)
	util.GetHttpResponse(w, r, user, err, true) 
}
