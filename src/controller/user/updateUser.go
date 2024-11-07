package user

import (
	"net/http"

	entities "domain/entity"
	"usecase/user"
	"util"
)

type UpdateUserController struct {
	Usecase user.UpdateUserUsecase
}

func (controller UpdateUserController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body entities.User
	err1 := util.GetRequestBody(r, &body)
	user, err2 := controller.Usecase.UpdateUser(body.Id, &body)

	var err error
	if err1 == nil {
		err = err2
	}

	util.GetHttpResponse(w, r, user, err, true)
}
