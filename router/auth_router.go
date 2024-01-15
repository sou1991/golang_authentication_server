package router

import (
	"github.com/sou1991/golang_authentication_server/controller"
)

func NewAuthRouter(ctrl controller.Authenticator) authRouter {
	return authRouter{ca: ctrl}
}

type authRouter struct {
	ca controller.Authenticator
}
