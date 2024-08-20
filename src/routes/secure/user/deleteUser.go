package user

import (
	"net/http"
	"strconv"
	"util"
	"controller/user"
)

const (
	deleteUserPattern string = "/user/{id}"
	deleteUserMethod string = http.MethodDelete
)

type DeleteUserRoute struct {
	Controller user.DeleteUserController
}

func (route DeleteUserRoute) Pattern() string {
	return deleteUserPattern
}

func (route DeleteUserRoute) Method() string {
	return deleteUserMethod
}

func (route DeleteUserRoute) HandleRequest(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := route.Controller.DeleteUser(id)
	util.GetHttpResponse(w, r, user, err, true)
}
