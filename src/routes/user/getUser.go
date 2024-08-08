package user

import (
	"domain"
	"net/http"
	"strconv"
	controller "controller/user"
	"util"
)


func getUserHandler(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := controller.GetUser(id);
	util.GetHttpResponse(w, r, user, err)
}

var GetUser domain.Route = domain.Route{
	Pattern: "/user/{id}",
	Method: http.MethodGet,
	Handler: getUserHandler,
}