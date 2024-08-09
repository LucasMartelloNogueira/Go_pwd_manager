package auth

import (
	"domain"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}


var Login domain.Route = domain.Route{
	Pattern: "/auth",
	Method: http.MethodPost,
	Handler: loginHandler,
}