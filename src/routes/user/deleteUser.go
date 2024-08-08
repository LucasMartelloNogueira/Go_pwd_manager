package user

import (
	"domain"
	"net/http"
	"strconv"
	controller "controller/user"
	"util"
)

func deleteUserHandler(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := controller.DeleteUser(id);
	util.GetHttpResponse(w, r, user, err)
}

var DeleteUser domain.Route = domain.Route{
	Pattern: "/user/{id}",
	Method: http.MethodDelete,
	Handler: deleteUserHandler,
}