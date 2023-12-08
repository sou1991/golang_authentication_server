package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/controller"
	"github.com/sou1991/golang_authentication_server/repository"
)

var rc = repository.NewClientRepository()
var cc = controller.NewClentController(rc)
var cro = NewClientRouter(cc)

var ra = repository.NewAuthRepository()
var ca = controller.NewAuthController(ra)
var aro = NewAuthRouter(ca)

func Router(r *gin.Engine) {
	r.GET("/outh2/v1/auth", func(c *gin.Context) {
		cro.Cc.CheckClent(c)
	})

	r.POST("/outh2/v1/auth/login", func(c *gin.Context) {
		aro.Ca.Auth(c)
	})

	r.POST("/outh2/v1/auth/token", func(c *gin.Context) {
		aro.Ca.Access(c)
	})
}
