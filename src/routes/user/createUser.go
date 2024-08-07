package user

import (
	"domain"
	"net/http"
	"util"
	controller "controller/user"
	"encoding/json"
	"fmt"
)


func createUserHandler(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-type", "application/json")
	var body domain.CreateUserBody;
	util.GetRequestBody(r, &body)

	user, err := controller.CreateUser(&body);

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

var CreateUser domain.Route = domain.Route{
	Pattern: "/user",
	Method: http.MethodPost,
	Handler: createUserHandler,
}