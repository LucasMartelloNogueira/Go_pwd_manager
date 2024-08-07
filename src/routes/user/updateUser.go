package user

import (
	"domain"
	"net/http"
	"util"
	"encoding/json"
	"fmt"
	controller "controller/user"

)

func updateUserHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	var body domain.User;
	util.GetRequestBody(r, &body)

	user, err := controller.UpdateUser(&body);

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

var UpdateUser domain.Route = domain.Route{
	Pattern: "/user",
	Method: http.MethodPut,
	Handler: updateUserHandler,
}