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
	var newUser entities.NewUser
	util.GetRequestBody(r, &newUser)

	userWithId, err := controller.Usecase.Register(&newUser)
	util.GetHttpResponse(w, r, userWithId, err, false)
}
