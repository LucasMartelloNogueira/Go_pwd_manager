package greet

import (
	"fmt"
	"net/http"
	"domain"
	"encoding/json"
	"util"
)

func greetRouteHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Header().Add("Content-type", "application/json")

	data := map[string]any{
		"message": fmt.Sprintf("Hello, %s", name),
	}

	body := map[string]any{
		"statusCode": http.StatusOK,
		"data": data,
	}

	bodyBytes, err := json.MarshalIndent(body, "", "  ")
	util.CheckErr(err, "[GreetRoute] error encoding response json")
	fmt.Fprintf(w, string(bodyBytes));
}


var GreetRoute domain.Route = domain.Route{
	Pattern: "/greet/{name}",
	Method: http.MethodGet,
	Handler: greetRouteHandler,
}