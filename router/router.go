package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sou1991/golang_authentication_server/repository"
	"github.com/sou1991/golang_authentication_server/controller"
	"github.com/sou1991/golang_authentication_server/router"
)

var cr = repository.NewClientRepository()
var cc = controller.NewClentController(cr)
var cro = router.NewClientRouter(cc)

var ar = repository.NewAuthRepository()
var ac = controller.NewAuthController(ar)
var aro = router.NewAuthRouter(ac)

func Router(r *gin.Engine){
	r.GET("/outh2/v1/auth", func(c *gin.Context) {
		cr.Cc.CheckClent(c)
	})

	r.POST("/outh2/v1/auth/login", func(c *gin.Context) {
		ar.Ca.Auth(c)
	})
}
