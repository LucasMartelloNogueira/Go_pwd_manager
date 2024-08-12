package user

import (
	controller "controller/user"
	types "domain/types"
	"net/http"
	"strconv"
	"util"
)

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := controller.DeleteUser(id)
	util.GetHttpResponse(w, r, user, err, true)
}

var DeleteUser types.Route = types.Route{
	Pattern: "/user/{id}",
	Method:  http.MethodDelete,
	Handler: deleteUserHandler,
}
