package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/controller"
)

type ClientChecker interface{
	CheckClent(...*gin.Engine)
}

type ClientRouter struct{
	Cc controller.ClientChecker
}

func NewClientRouter(cc controller.ClientChecker) ClientChecker{
	return ClientRouter{Cc: cc}
}

func (cr ClientRouter) CheckClent(r ...*gin.Engine){
	r[0].GET("/outh2/v1/auth", func(c *gin.Context) {
		cr.Cc.CheckClent(c)
	})
	
}