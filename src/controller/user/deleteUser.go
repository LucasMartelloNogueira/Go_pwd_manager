package user

import (
	"net/http"
	"strconv"
	"usecase/user"
	"util"
)

type DeleteUserController struct {
	Usecase user.DeleteUserUsecase
}

func (controller DeleteUserController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := controller.Usecase.DeleteUser(id)
	util.GetHttpResponse(w, r, user, err, true)
}
