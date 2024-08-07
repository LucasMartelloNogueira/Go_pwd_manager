package user

import (
	"domain"
	"net/http"
	"strconv"
	"fmt"
	controller "controller/user"
	"encoding/json"
	"util"
)

func deleteUserHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	id, _ := strconv.Atoi(r.PathValue("id"))
	
	user, err := controller.DeleteUser(id);

	if err != nil {
		errorResponse := map[string]any{
			"statusCode": 500,
			"error": "Internal server error",
		}
		responseInBytes, _ := json.MarshalIndent(errorResponse, "", "  ")
		fmt.Fprintf(w, string(responseInBytes))
		return
	}

	responseInBytes, err := json.MarshalIndent(user, "", "  ")
	util.CheckErr(err, "[GreetRoute] error encoding response json")
	fmt.Fprintf(w, string(responseInBytes))
}

var DeleteUser domain.Route = domain.Route{
	Pattern: "/user/{id}",
	Method: http.MethodDelete,
	Handler: deleteUserHandler,
}