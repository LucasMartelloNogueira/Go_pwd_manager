package password

import (
	"domain"
	"net/http"
	"encoding/json"
	"io"
	"util"
)

type getPasswordBody struct {
	userId int
	masterPassword string
}

func getPasswordHandler(w http.ResponseWriter, r *http.Request){
	var body getPasswordBody;

	bodyBytes, err := io.ReadAll(r.Body)
	util.CheckErr(err, "[GetPasswordRoute] io error reading request body")

	err = json.Unmarshal(bodyBytes, &body)
	util.CheckErr(err, "[GetPasswordRoute] error reading body from request")
}

var GetPassword domain.Route = domain.Route{
	Pattern: "/getPassword",
	Method: http.MethodPost,
	Handler: getPasswordHandler,
}