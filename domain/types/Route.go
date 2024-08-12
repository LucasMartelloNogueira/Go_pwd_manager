package domain

import (
	"fmt"
	"net/http"
)

type Route struct {
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

func (r *Route) GetMethodPatternStr() string {
	return fmt.Sprintf("%s %s", r.Method, r.Pattern)
}
