package user

import (
	"net/http"
	"strconv"

	"usecase/user"
	"util"
)

type GetUserController struct {
	Usecase user.GetUserUsecase
}

func (controller GetUserController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := controller.Usecase.GetUser(id)
	util.GetHttpResponse(w, r, user, err, true)
}
