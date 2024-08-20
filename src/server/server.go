package server

import (
	"log"
	"net/http"
	"routes"
	types "domain/types"
	"fmt"
)

func getMethodPatternStr(route types.Route) string{
	return fmt.Sprintf("%s %s", route.Method(), route.Pattern())
}

func InitServer() {
	for _, route := range routes.Routes {
		http.HandleFunc(getMethodPatternStr(route), route.HandleRequest)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
