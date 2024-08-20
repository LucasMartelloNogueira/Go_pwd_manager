package helloWorld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"util"
)

const (
	helloWorldPattern = "/helloWorld"
	helloWorldMethod = http.MethodGet
)

type HelloWorldRoute struct {}

func (route HelloWorldRoute) Pattern() string {
	return helloWorldPattern
}

func (route HelloWorldRoute) Method() string {
	return helloWorldMethod
}

func (route HelloWorldRoute) HandleRequest(w http.ResponseWriter, r *http.Request) {
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
