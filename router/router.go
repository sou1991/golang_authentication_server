package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/controller"
	"github.com/sou1991/golang_authentication_server/repository"
)

var cr = repository.NewClientRepository()
var cc = controller.NewClentController(cr)
var cro = NewClientRouter(cc)

var ar = repository.NewAuthRepository()
var ac = controller.NewAuthController(ar)
var aro = NewAuthRouter(ac)

func Router(r *gin.Engine) {
	r.GET("/outh2/v1/auth", func(c *gin.Context) {
		cro.Cc.CheckClent(c)
	})

	r.POST("/outh2/v1/auth/login", func(c *gin.Context) {
		aro.Ca.Auth(c)
	})
}
