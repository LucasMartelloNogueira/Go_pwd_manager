package user

import (
	entities "domain/entity"
	"net/http"
	usecase "usecase/user"
	"util"
)

type CreateUserController struct {
	Usecase usecase.CreateUserUsecase
}

func (controller CreateUserController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body entities.NewUser
	err1 := util.GetRequestBody(r, &body)
	user, err2 := controller.Usecase.CreateUser(&body)

	var err error
	if err1 == nil {
		err = err2
	}

	util.GetHttpResponse(w, r, user, err, true)
}
