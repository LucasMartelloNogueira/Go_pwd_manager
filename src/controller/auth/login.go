package auth

import (
	dtos "domain/dto"
	"net/http"
	"usecase/auth"
	"util"
)

type LoginController struct {
	Usecase auth.LoginUsecase
}

func (controller LoginController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var userCredentials dtos.UserCredentials
	err1 := util.GetRequestBody(r, &userCredentials)
	loginInfo, err2 := controller.Usecase.Login(userCredentials)

	var err error
	if err1 == nil {
		err = err2
	}

	util.GetHttpResponse(w, r, loginInfo, err, false)
}
