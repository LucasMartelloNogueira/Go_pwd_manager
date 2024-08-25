package auth

import (
	entities "domain/entity"
	"net/http"
	"usecase/auth"
	"util"
)

type RegisterController struct {
	Usecase auth.RegisterUsecase
}

func (controller RegisterController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	util.GetRequestBody(r, &user)

	userWithId, err := controller.Usecase.Register(&user)
	util.GetHttpResponse(w, r, userWithId, err, false)
}
