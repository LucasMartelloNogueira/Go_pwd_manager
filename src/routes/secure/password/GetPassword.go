package password

import (
	types "domain/types"
	"net/http"
)

func getPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

var GetPassword types.Route = types.Route{
	Pattern: "/getPassword",
	Method:  http.MethodPost,
	Handler: getPasswordHandler,
}
