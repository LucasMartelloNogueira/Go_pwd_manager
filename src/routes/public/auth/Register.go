package auth

import (
	"domain"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}


var Register domain.Route = domain.Route{
	Pattern: "/auth",
	Method: http.MethodPost,
	Handler: registerHandler,
}