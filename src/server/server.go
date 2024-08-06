package server

import (
	"log"
	"net/http"
	"routes"
)

func InitServer() {
	for _, route := range routes.Routes {
		http.HandleFunc(route.GetMethodPatternStr(), route.Handler)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}