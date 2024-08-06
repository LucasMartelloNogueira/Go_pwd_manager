package userRoute

import (
	"domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"usecaseUser"
	"util"
)


func getUserHandler(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-type", "application/json")
	id, _ := strconv.Atoi(r.PathValue("id"))
	
	user, err := usecaseUser.GetUser(id);

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

var GetUserRoute domain.Route = domain.Route{
	Pattern: "/user/{id}",
	Method: http.MethodGet,
	Handler: getUserHandler,
}