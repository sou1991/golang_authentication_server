package router

import (
	"github.com/sou1991/golang_authentication_server/controller"
)

type clientRouter struct {
	cc controller.ClientChecker
}

func NewClientRouter(ctrl controller.ClientChecker) clientRouter {
	return clientRouter{cc: ctrl}
}
