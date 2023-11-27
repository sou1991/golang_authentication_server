package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/controller"
)

func NewAuthRouter(ca controller.Authenticator) Authenticator{
	return authRouter{Ca: ca}
}

type authRouter struct{
	Ca controller.Authenticator
}

type Authenticator interface{
	Auth(...*gin.Engine)
}

func (ar authRouter) Auth(r ...*gin.Engine){
	r[0].GET("/outh2/v1/auth/login", func(c *gin.Context) {
		ar.Ca.Auth(c)
	})
}