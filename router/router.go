package router

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/sou1991/golang_authentication_server/controller"
	repo "github.com/sou1991/golang_authentication_server/repository"
)

var rc repo.ClientChecker = repo.NewClientRepository()
var cc ctrl.ClientChecker= ctrl.NewClentController(rc)
var cro = NewClientRouter(cc)

var ra repo.Authenticator = repo.NewAuthRepository()
var ca ctrl.Authenticator = ctrl.NewAuthController(ra)
var aro = NewAuthRouter(ca)

func Router(r *gin.Engine) {
	r.GET("/outh2/v1/auth", func(c *gin.Context) {
		cro.cc.CheckClent(c)
	})

	r.POST("/outh2/v1/auth/login", func(c *gin.Context) {
		aro.ca.Auth(c)
	})

	r.POST("/outh2/v1/auth/token", func(c *gin.Context) {
		aro.ca.Access(c)
	})
}
