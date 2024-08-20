package domain

import (
	"net/http"
)

type Route interface {
	Pattern() string
	Method() string
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

// type Route struct {
// 	Pattern string
// 	Method  string
// 	Handler http.HandlerFunc
// }

// func (r *Route) GetMethodPatternStr() string {
// 	return fmt.Sprintf("%s %s", r.Method, r.Pattern)
// }
