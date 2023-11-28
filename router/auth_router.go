package router

import (
	"github.com/sou1991/golang_authentication_server/controller"
)

func NewAuthRouter(ca controller.Authenticator) authRouter {
	return authRouter{Ca: ca}
}

type authRouter struct {
	Ca controller.Authenticator
}
