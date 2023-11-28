package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/controller"
)

type ClientRouter struct{
	Cc controller.ClientChecker
}

func NewClientRouter(cc controller.ClientChecker) ClientChecker{
	return ClientRouter{Cc: cc}
}
