package routes

import (
	"domain"
	"greet"
	"userRoute"
)

var Routes []domain.Route = []domain.Route{
	greet.GreetRoute,
	userRoute.CreateUserRoute,
	userRoute.GetUserRoute,
}

