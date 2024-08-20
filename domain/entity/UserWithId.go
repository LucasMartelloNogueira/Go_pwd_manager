package domain

type UserWithId struct {
	Id       int `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


type UserWithIdOpt func(userWithId *UserWithId)

func WithId(id int) UserWithIdOpt {
	return func(userWithId *UserWithId) {
		userWithId.Id = id
	}
}

func WithName(name string) UserWithIdOpt {
	return func(userWithId *UserWithId) {
		userWithId.Name = name
	}
}

func WithEmail(email string) UserWithIdOpt {
	return func(userWithId *UserWithId) {
		userWithId.Email = email
	}
}

func WithPassWord(password string) UserWithIdOpt {
	return func(userWithId *UserWithId) {
		userWithId.Password = password
	}
}

func UserWithIdBuilder(opts ...UserWithIdOpt) UserWithId {
	userWithId := UserWithId{
		Id: 0,
		Name: "",
		Email: "",
		Password: "",
	}

	for _, opt := range opts {
		opt(&userWithId)
	}

	return userWithId
}
