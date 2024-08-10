package user

import (
	controller "controller/user"
	types "domain/types"
	"net/http"
	"strconv"
	"util"
)


func getUserHandler(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := controller.GetUser(id)
	util.GetHttpResponse(w, r, user, err, true)
}

var GetUser types.Route = types.Route{
	Pattern: "/user/{id}",
	Method: http.MethodGet,
	Handler: getUserHandler,
}