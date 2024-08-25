package domain

type LoginResponse struct {
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}

type LoginResponseOpt func(r *LoginResponse)

func WithUserId(userId int) LoginResponseOpt {
	return func(r *LoginResponse) {
		r.UserId = userId
	}
}

func WithToken(token string) LoginResponseOpt {
	return func(r *LoginResponse) {
		r.Token = token
	}
}

func LoginResponseBuilder(opts ...LoginResponseOpt) LoginResponse {

	var loginResponse LoginResponse = LoginResponse{
		UserId: 0,
		Token:  "",
	}

	for _, opt := range opts {
		opt(&loginResponse)
	}

	return loginResponse
}
