package domain

import (
	"net/http"
)

type Controller interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}
