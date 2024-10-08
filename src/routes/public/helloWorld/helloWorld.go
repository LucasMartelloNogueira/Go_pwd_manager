package helloWorld

import (
	types "domain/types"
	"encoding/json"
	"fmt"
	"net/http"
	"util"
)

type HelloWorldController struct{}

func (controller HelloWorldController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	body := map[string]any{
		"data": map[string]any{
			"message": "Hello, World!",
		},
		"status":     util.HttpStatusSuccess,
		"statusCode": http.StatusOK,
	}

	bodyBytes, _ := json.MarshalIndent(body, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, string(bodyBytes))
}

var HelloWorldRoute types.Route = types.Route{
	Pattern:    "/helloWorld",
	Method:     http.MethodGet,
	Controller: HelloWorldController{},
}
