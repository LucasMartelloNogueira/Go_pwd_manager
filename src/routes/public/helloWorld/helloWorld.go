package helloWorld

import (
	"fmt"
	"net/http"
	"domain"
	"encoding/json"
	"util"
)

func helloWorlHandler(w http.ResponseWriter, r *http.Request) {
	body := map[string]any{
		"data": map[string]any{
			"message": "Hello, World!",
		},
		"status": util.HttpStatusSuccess,
		"statusCode": http.StatusOK,
	}

	bodyBytes, _ := json.MarshalIndent(body, "", "  ")
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	fmt.Fprint(w, string(bodyBytes));
}


var HelloWorld domain.Route = domain.Route{
	Pattern: "/helloWorld",
	Method: http.MethodGet,
	Handler: helloWorlHandler,
}