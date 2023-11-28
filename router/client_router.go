package router

import (
	"github.com/sou1991/golang_authentication_server/controller"
)

type clientRouter struct {
	Cc controller.ClientChecker
}

func NewClientRouter(cc controller.ClientChecker) clientRouter {
	return clientRouter{Cc: cc}
}
