package domain

import (
	"fmt"
)

type Route struct {
	Pattern    string
	Method     string
	Controller Controller
}

func (r *Route) GetMethodPatternStr() string {
	return fmt.Sprintf("%s %s", r.Method, r.Pattern)
}
